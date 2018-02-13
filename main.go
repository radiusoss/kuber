package main

import (
	"fmt"
	"time"

	"github.com/datalayer/kuber/aws"
	"github.com/datalayer/kuber/cmd"
)

func main() {
	go sanitize()
	cmd.Execute()
}

func sanitize() {
	region := "eu-central-1"
	for true {
		tagKuberWorkers(region)
		registerMasterToLoadBalancers(region)
		time.Sleep(time.Second * time.Duration(30))
	}
}

func tagKuberWorkers(region string) {
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
	spitfireLb := aws.GetLoadBalancersByTag("kuber-role", "spitfire", region)[0]
	fmt.Println("Spitfire Load Balancer: " + *spitfireLb)
	spitfireResult := aws.RegisterInstanceToLoadBalancer(inst, spitfireLb, region)
	fmt.Println(spitfireResult)
	explorerLb := aws.GetLoadBalancersByTag("kuber-role", "explorer", region)[0]
	fmt.Println("Explorer Load Balancer: " + *explorerLb)
	explorerResult := aws.RegisterInstanceToLoadBalancer(inst, explorerLb, region)
	fmt.Println(explorerResult)
}
