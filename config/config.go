package config

const DefaultRegion = "us-west-2"

var KuberConfig Config

type Config struct {
	KuberUi                string `json:"kuberUi"`
	KuberRest              string `json:"kuberRest"`
	KuberWs                string `json:"kuberWs"`
	SpitfireRest           string `json:"spitfireRest"`
	SpitfireWs             string `json:"spitfireWs"`
	//	Server   ServerConfig
	//	Database PersistenceConfig
}

type ServerConfig struct {
	Port int
}

type PersistenceConfig struct {
	ConnectionUri string
}
