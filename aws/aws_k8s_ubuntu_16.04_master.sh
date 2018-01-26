# ------------------------------------------------------------------------------------------------------------------------
# We are explicitly not using a templating language to inject the values as to encourage the user to limit their
# use of templating logic in these files. By design all injected values should be able to be set at runtime,
# and the shell script real work. If you need conditional logic, write it in bash or make another shell script.
# ------------------------------------------------------------------------------------------------------------------------

apt install -y awscli

INSTANCEID=`/usr/bin/curl -s http://169.254.169.254/latest/meta-data/instance-id`
echo $INSTANCEID
REGION=`curl http://169.254.169.254/latest/dynamic/instance-identity/document | grep region | awk -F\" '{print $4}'`
echo $REGION

OUTPUT="text"
# Check if instance has a public IP from Elastic pool assigned
PUBLICIPASSIGNED=`aws ec2 describe-addresses --output ${OUTPUT} --region ${REGION} | grep ${INSTANCEID} | wc -l`
if [ ${PUBLICIPASSIGNED} -gt 0 ]
then
  # Instance has (at least) one Public IP associated from Elastic pool.
  IP=`aws ec2 describe-addresses --output ${OUTPUT} --region ${REGION} | grep ${INSTANCEID} | head -1 | awk '{ print $NF; }'`
  echo "Alreay running with IP from EIP pool: "$IP
else
  # Get IP address from Elastic pool
  IP=`aws ec2 describe-addresses --output ${OUTPUT} --region ${REGION} | grep -v 'i-' | head -1 | awk '{ print $NF; }'`
  echo "Found IP in EIP pool: "$IP
  if [ "${IP}" = "" ]
  then
    # Public IP on Elastic pool unavailable...
    echo "Free Public IP inside Elastic pool in ${REGION} region unavailable"
  else
    # Attach Public IP from Elastic pool
    echo "Attaching IP from EIP pool: "$IP
    aws ec2 associate-address --output ${OUTPUT} --region ${REGION} --instance-id ${INSTANCEID} --public-ip ${IP}
  fi
fi

aws ec2 create-tags --resources ${INSTANCEID} --region ${REGION} --tags Key=Cost,Value=kuber

# Specify the Kubernetes version to use
KUBERNETES_VERSION="1.8.6"

curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
touch /etc/apt/sources.list.d/kubernetes.list
sh -c 'echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list'

# Has to be configured before installing kubelet, or kubelet has to be restarted to pick up changes
mkdir -p /etc/systemd/system/kubelet.service.d
touch /etc/systemd/system/kubelet.service.d/20-cloud-provider.conf
cat << EOF  > /etc/systemd/system/kubelet.service.d/20-cloud-provider.conf
[Service]
Environment="KUBELET_EXTRA_ARGS=--cloud-provider=aws --runtime-request-timeout 4m0s"
EOF

chmod 0600 /etc/systemd/system/kubelet.service.d/20-cloud-provider.conf

apt-get update -y
apt-get install -y \
    socat \
    ebtables \
    docker.io \
    apt-transport-https \
    kubelet \
    kubeadm=${KUBERNETES_VERSION}-00 \
    cloud-utils \
    jq


systemctl enable docker
systemctl start docker

PUBLICIP=$(ec2metadata --public-ipv4 | cut -d " " -f 2)
PRIVATEIP=$(ec2metadata --local-ipv4 | cut -d " " -f 2)
TOKEN=$(cat /etc/kubicorn/cluster.json | jq -r '.values.itemMap.INJECTEDTOKEN')
PORT=$(cat /etc/kubicorn/cluster.json | jq -r '.values.itemMap.INJECTEDPORT | tonumber')

# Necessary for joining a cluster with AWS information
HOSTNAME=$(hostname -f)

cat << EOF  > "/etc/kubicorn/kubeadm-config.yaml"
apiVersion: kubeadm.k8s.io/v1alpha1
kind: MasterConfiguration
cloudProvider: aws
token: ${TOKEN}
kubernetesVersion: ${KUBERNETES_VERSION}
nodeName: ${HOSTNAME}
api:
  advertiseAddress: ${PUBLICIP}
  bindPort: ${PORT}
apiServerCertSANs:
- ${PUBLICIP}
- ${HOSTNAME}
- ${PRIVATEIP}
authorizationModes:
- Node
- RBAC
EOF

kubeadm reset
kubeadm init --config /etc/kubicorn/kubeadm-config.yaml

# Thanks Kelsey :)
kubectl apply \
  -f http://docs.projectcalico.org/v2.3/getting-started/kubernetes/installation/hosted/kubeadm/1.6/calico.yaml \
  --kubeconfig /etc/kubernetes/admin.conf

kubectl apply \
    -f  https://raw.githubusercontent.com/kubernetes/kubernetes/release-1.8/cluster/addons/storage-class/aws/default.yaml \
    --kubeconfig /etc/kubernetes/admin.conf

mkdir -p /home/ubuntu/.kube
cp /etc/kubernetes/admin.conf /home/ubuntu/.kube/config
chown -R ubuntu:ubuntu /home/ubuntu/.kube

mkdir -p ~/.kube
cp /etc/kubernetes/admin.conf ~/.kube/config

alias k=kubectl

kubectl label nodes ${HOSTNAME} kuber-role=master
# Taint master to host pods.
# This will remove the node-role.kubernetes.io/master taint from any nodes that have it, including the master node, 
# meaning that the scheduler will then be able to schedule pods everywhere.
kubectl taint nodes --all node-role.kubernetes.io/master-

function setup_rbac() {
  cat << EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: dla-rbac
subjects:
  - kind: ServiceAccount
    # Reference to upper's `metadata.name`
    name: default
    # Reference to upper's `metadata.namespace`
    namespace: default
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
EOF
  kubectl create clusterrolebinding add-on-cluster-admin \
    --clusterrole=cluster-admin \
    --serviceaccount=kube-system:default
}

function install_helm() {
  wget https://storage.googleapis.com/kubernetes-helm/helm-v2.7.2-linux-amd64.tar.gz
  tar xvfz helm-v2.7.2-linux-amd64.tar.gz
  mv linux-amd64/helm /usr/local/bin
  helm init --canary-image --upgrade; kubectl rollout status -w deployment/tiller-deploy --namespace=kube-system;
  kubectl create serviceaccount --namespace kube-system tiller
  kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
  kubectl patch deploy --namespace kube-system tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'      
  helm init --service-account tiller --upgrade
  kubectl rollout status -w deployment/tiller-deploy --namespace=kube-system
}

setup_rbac
install_helm

# git clone https://github.com/datalayer/helm-charts
# cd helm-charts
# ./deploy-chart.sh heapster
# ./deploy-chart.sh dashboard

# helm install -n heapster \
#   --namespace kube-system \
#   stable/heapster

# helm install stable/kubernetes-dashboard \
#   --namespace kube-system \
#   --set=httpPort=3000,resources.limits.cpu=200m,rbac.create=true \
#   -n k8s-dashboard
