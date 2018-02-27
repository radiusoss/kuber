package ws

import (
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/datalayer/kuber/slots"
	corev1 "k8s.io/api/core/v1"
)

type WsMessage struct {
	Op      string        `json:"op"`
	Message interface{}   `json:"message"`
	Cluster ClusterStatus `json:"cluster,omitempty"`
	Slots   []slots.Slot  `json:"slots,omitempty"`
}

type ClusterStatus struct {
	ClusterName         string                       `json:"clusterName,omitempty"`
	AwsProfile          string                       `json:"awsProfile,omitempty`
	AwsInstances        *ec2.DescribeInstancesOutput `json:"awsInstances,omitempty"`
	AwsAutoscalingGroup *autoscaling.Group           `json:"awsAutoscalingGroup,omitempty"`
	Nodes               *corev1.NodeList             `json:"nodes,omitempty"`
	Pods                *corev1.PodList              `json:"pods,omitempty"`
}
