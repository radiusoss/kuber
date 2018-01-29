package config

var KuberConfig Config

type Config struct {
	Hdfs                   string `json:"hdfs"`
	KuberBoard             string `json:"kuberBoard"`
	KuberRest              string `json:"kuberRest"`
	KuberWs                string `json:"kuberWs"`
	MicrosoftApplicationId string `json:"microsoftApplicationId"`
	MicrosoftRedirect      string `json:"microsoftRedirect"`
	MicrosoftScope         string `json:"microsoftScope"`
	MicrosoftSecret        string `json:"microsoftSecret"`
	SpitfireRest           string `json:"spitfireRest"`
	SpitfireWs             string `json:"spitfireWs"`
	TwitterConsumerKey     string `json:"consumerKey"`
	TwitterConsumerSecret  string `json:"consumerSecret"`
	TwitterRedirect        string `json:"twitterRedirect"`
	//	Server   ServerConfig
	//	Database PersistenceConfig
}

type ServerConfig struct {
	Port int
}

type PersistenceConfig struct {
	ConnectionUri string
}
