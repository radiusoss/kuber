package azure

type AzureKuberCluster struct {
	KuberClusterId    uint
	ResourceGroup     string
	AgentCount        int
	AgentName         string
	KubernetesVersion string
}

type CreateAzureCluster struct {
	Node *CreateAzureNode `json:"node"`
}

type CreateAzureNode struct {
	ResourceGroup     string `json:"resourceGroup"`
	AgentCount        int    `json:"agentCount"`
	AgentName         string `json:"agentName"`
	KubernetesVersion string `json:"kubernetesVersion"`
}

type UpdateAzureCluster struct {
	*UpdateAzureNode `json:"node"`
}

type UpdateAzureNode struct {
	AgentCount int `json:"agentCount"`
}
