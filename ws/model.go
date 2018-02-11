package ws

import "github.com/datalayer/kuber/slots"

type WsMessage struct {
	Op      string         `json:"op"`
	Message interface{}    `json:"message"`
	Cluster ClusterContent `json:"cluster,omitempty"`
	Slots   []slots.Slot   `json:"slots,omitempty"`
}

type ClusterContent struct {
	ClusterName string `json:"clusterName,omitempty"`
	AwsProfile  string `json:"aswProfile,omitempty`
}
