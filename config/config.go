package config

var KuberConfig Config

type Config struct {
	KuberRest          string `json:"kuberRest"`
	KuberWs            string `json:"kuberWs"`
	AzureApplicationId string `json:"azureApplicationId"`
	AzureRedirect      string `json:"azureRedirect"`
	AzureScope         string `json:"azureScope"`
	SpitfireRest       string `json:"spitfireRest"`
	SpitfireWs         string `json:"spitfireWs"`
	Hdfs               string `json:"hdfs"`
	TwitterRedirect    string `json:"twitterRedirect"`
	KuberPlane         string `json:"kuberPlane"`
	//	Server   ServerConfig
	//	Database PersistenceConfig
}

type ServerConfig struct {
	Port int
}

type PersistenceConfig struct {
	ConnectionUri string
}
