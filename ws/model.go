package ws

type WsMessage struct {
	Op      string         `json:"op"`
	Message interface{}    `json:"message"`
	Cluster ClusterContent `json:"cluster,omitempty"`
}

type ClusterContent struct {
	ClusterName string `json:"clusterName,omitempty"`
	AwsProfile  string `json:"aswProfile,omitempty`
}
