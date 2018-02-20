package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/datalayer/kuber/aws"
	"github.com/datalayer/kuber/cmd"
	"github.com/datalayer/kuber/slots"
	"github.com/datalayer/kuber/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	go sanitize()
	cmd.Execute()
}

func sanitize() {
	region := "eu-central-1"

	var config *rest.Config

	if _, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount/token"); os.IsNotExist(err) {
		var kubeconfig *string
		if home := util.GetUserHome(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}

	} else {

		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}

	}

	for os.Getenv("KUBER_AUTOSCALING") == "true" {
		adjustNodeCapacity(region)
		tagK8SWorkers(config, region)
		registerMasterToLoadBalancers(region)
		time.Sleep(time.Second * time.Duration(15))
	}
}

func adjustNodeCapacity(region string) {
	slt := slots.Slots
	clusterUp := false
	for _, s := range slt {
		now := time.Now()
		start, _ := time.Parse(time.RFC3339, s.Start)
		end, _ := time.Parse(time.RFC3339, s.End)
		fmt.Println("Slot %v %v", start, end)
		if now.After(start) && now.Before(end) {
			clusterUp = true
		}
	}
	var desiredWorkers int64 = 0
	if clusterUp {
		desiredWorkers = 3
	}
	fmt.Println("+++ Desired Workers: %v", desiredWorkers)
	aws.ScaleWorkers(desiredWorkers, region)
}

func tagK8SWorkers(config *rest.Config, region string) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, node := range nodes.Items {
		if node.Labels["kuber-role"] != "master" {
			fmt.Println(node.Name)
			l := node.Labels
			l["kuber-role"] = "node"
			node.SetLabels(l)
			clientset.CoreV1().Nodes().Update(&node)
		}
	}

}

func tagAwsWorkers(region string) {
	resp := aws.InstancesByTag("Name", "kuber.node", region)
	if resp.Reservations != nil {
		for _, instance := range resp.Reservations[0].Instances {
			id := *instance.InstanceId
			fmt.Println("Tagging with kuber-role=node resource: " + id)
			aws.TagResource(id, "kuber-role", "node", region)
		}
	}
}

func registerMasterToLoadBalancers(region string) {
	inst := aws.InstancesByTag("Name", "kuber.master", region).Reservations[0].Instances[0].InstanceId
	fmt.Println("Master Instance: " + *inst)
	spitfireLbs := aws.GetLoadBalancersByTag("kuber-role", "spitfire", region)
	if len(spitfireLbs) > 0 {
		spitfireLb := spitfireLbs[0]
		fmt.Println("Spitfire Load Balancer: " + *spitfireLb)
		spitfireResult := aws.RegisterInstanceToLoadBalancer(inst, spitfireLb, region)
		fmt.Println(spitfireResult)
	}
	explorerLbs := aws.GetLoadBalancersByTag("kuber-role", "explorer", region)
	if len(explorerLbs) > 0 {
		explorerLb := explorerLbs[0]
		fmt.Println("Explorer Load Balancer: " + *explorerLb)
		explorerResult := aws.RegisterInstanceToLoadBalancer(inst, explorerLb, region)
		fmt.Println(explorerResult)
	}
}
