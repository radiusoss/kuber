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

aws ec2 create-tags --resources ${INSTANCEID} --region ${REGION} --tags Key=Cost,Value=kuber

rm -fr /var/lib/docker
mkdir /mnt/docker
ln -s /mnt/docker /var/lib/docker

# Specify the Kubernetes version to use
KUBERNETES_VERSION="1.8.6"

curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
touch /etc/apt/sources.list.d/kubernetes.list
sh -c 'echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list'

# Has to be configured before installing kubelet, or kubelet has to be restarted to pick up changes
mkdir -p /etc/systemd/system/kubelet.service.d
touch /etc/systemd/system/kubelet.service.d/20-cloud-provider.conf
# Do we need to tune the docker pull timeout? It does not seem to give good results...
# --runtime-request-timeout 4m0s
cat << EOF  > /etc/systemd/system/kubelet.service.d/20-cloud-provider.conf
[Service]
Environment="KUBELET_EXTRA_ARGS=--cloud-provider=aws"
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
    jq

systemctl enable docker
systemctl start docker

TOKEN=$(cat /etc/kubicorn/cluster.json | jq -r '.values.itemMap.INJECTEDTOKEN')
MASTER=$(cat /etc/kubicorn/cluster.json | jq -r '.values.itemMap.INJECTEDMASTER')

systemctl daemon-reload
systemctl restart kubelet.service

# Necessary for joining a cluster with the AWS information
HOSTNAME=$(hostname -f)

kubeadm reset

cat <<EOF > /root/kubeadm-aws-join.conf
apiVersion: kubeadm.k8s.io/v1alpha1
kind: NodeConfiguration
kubernetesVersion: ${KUBERNETES_VERSION}
token: ${TOKEN}
nodeName: ${HOSTNAME}
discoveryTokenAPIServers:
- ${MASTER}
EOF

kubeadm join --config=/root/kubeadm-aws-join.conf

# HOSTNAME=$(hostname -f)
# KUBECONFIG=/etc/kubernetes/kubelet.conf kubectl label nodes ${HOSTNAME} kuber-role=worker

# systemctl restart kubelet
reboot
