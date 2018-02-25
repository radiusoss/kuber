package cmd

import (
	"crypto/elliptic"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/datalayer/kuber/aws"
	"github.com/datalayer/kuber/config"
	"github.com/datalayer/kuber/handler"
	krest "github.com/datalayer/kuber/rest"
	"github.com/datalayer/kuber/slots"
	"github.com/datalayer/kuber/util"

	restful "github.com/emicklei/go-restful"
	"github.com/kubernetes/dashboard/src/app/backend/auth"
	authApi "github.com/kubernetes/dashboard/src/app/backend/auth/api"
	"github.com/kubernetes/dashboard/src/app/backend/auth/jwe"
	"github.com/kubernetes/dashboard/src/app/backend/cert"
	"github.com/kubernetes/dashboard/src/app/backend/cert/ecdsa"
	"github.com/kubernetes/dashboard/src/app/backend/client"
	clientapi "github.com/kubernetes/dashboard/src/app/backend/client/api"
	"github.com/kubernetes/dashboard/src/app/backend/integration"
	integrationapi "github.com/kubernetes/dashboard/src/app/backend/integration/api"
	"github.com/kubernetes/dashboard/src/app/backend/settings"
	"github.com/kubernetes/dashboard/src/app/backend/sync"
	"github.com/kubernetes/dashboard/src/app/backend/systembanner"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	argInsecurePort             int
	argPort                     int
	argInsecureBindAddress      net.IP
	argBindAddress              net.IP
	argDefaultCertDir           string
	argCertFile                 string
	argKeyFile                  string
	argApiserverHost            string
	argHeapsterHost             string
	argKubeConfigFile           string
	argTokenTTL                 int
	argAuthenticationMode       []string
	argMetricClientCheckPeriod  int
	argAutoGenerateCertificates bool
	argEnableInsecureLogin      bool
	argSystemBanner             string
	argSystemBannerSeverity     string
	argsKuberUi                 string
	argsKuberRest               string
	argsKuberWs                 string
	argsGoogleClientId          string
	argsGoogleRedirect          string
	argsGoogleSecret            string
	argsGoogleScope             string
	argsGoogleApiKey            string
	argsMicrosoftApplicationId  string
	argsMicrosoftRedirect       string
	argsMicrosoftSecret         string
	argsMicrosoftScope          string
	argsSpitfireRest            string
	argsSpitfireWs              string
	argsHdfs                    string
	argsTwitterConsumerKey      string
	argsTwitterConsumerSecret   string
	argsTwitterRedirect         string
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs the REST Server",
	Long:  `This subcommand runs the REST Server`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {

	KuberCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().IntVar(&argInsecurePort, "insecure-port", 9091, "The port to listen to for incoming HTTP requests.")
	serverCmd.PersistentFlags().IntVar(&argPort, "port", 8443, "The secure port to listen to for incoming HTTPS requests.")
	serverCmd.PersistentFlags().IPVar(&argInsecureBindAddress, "insecure-bind-address", net.IPv4(127, 0, 0, 1), "The IP address on which to serve the --port (set to 0.0.0.0 for all interfaces).")
	serverCmd.PersistentFlags().IPVar(&argBindAddress, "bind-address", net.IPv4(0, 0, 0, 0), "The IP address on which to serve the --secure-port (set to 0.0.0.0 for all interfaces).")
	serverCmd.PersistentFlags().StringVar(&argDefaultCertDir, "default-cert-dir", "/certs", "Directory path containing '--tls-cert-file' and '--tls-key-file' files. Used also when auto-generating certificates flag is set.")
	serverCmd.PersistentFlags().StringVar(&argCertFile, "tls-cert-file", "", "File containing the default x509 Certificate for HTTPS.")
	serverCmd.PersistentFlags().StringVar(&argKeyFile, "tls-key-file", "", "File containing the default x509 private key matching --tls-cert-file.")
	serverCmd.PersistentFlags().StringVar(&argApiserverHost, "apiserver-host", "", "The address of the Kubernetes Apiserver "+
		"to connect to in the format of protocol://address:port, e.g., "+
		"http://localhost:8080. If not specified, the assumption is that the binary runs inside a "+
		"Kubernetes cluster and local discovery is attempted.")
	serverCmd.PersistentFlags().StringVar(&argHeapsterHost, "heapster-host", "", "The address of the Heapster Apiserver "+
		"to connect to in the format of protocol://address:port, e.g., "+
		"http://localhost:8082. If not specified, the assumption is that the binary runs inside a "+
		"Kubernetes cluster and service proxy will be used.")
	serverCmd.PersistentFlags().StringVar(&argKubeConfigFile, "kubeconfig", "", "Path to kubeconfig file with authorization and master location information.")
	serverCmd.PersistentFlags().IntVar(&argTokenTTL, "token-ttl", int(authApi.DefaultTokenTTL), "Expiration time (in seconds) of JWE tokens generated by dashboard. Default: 15 min. 0 - never expires")
	serverCmd.PersistentFlags().StringSliceVar(&argAuthenticationMode, "authentication-mode", []string{authApi.Token.String()}, "Enables authentication options that will be reflected on login screen. Supported values: token, basic. Default: token."+
		"Note that basic option should only be used if apiserver has '--authorization-mode=ABAC' and '--basic-auth-file' flags set.")
	serverCmd.PersistentFlags().IntVar(&argMetricClientCheckPeriod, "metric-client-check-period", 30, "Time in seconds that defines how often configured metric client health check should be run. Default: 30 seconds.")
	serverCmd.PersistentFlags().BoolVar(&argAutoGenerateCertificates, "auto-generate-certificates", false, "When set to true, Dashboard will automatically generate certificates used to serve HTTPS. Default: false.")
	serverCmd.PersistentFlags().BoolVar(&argEnableInsecureLogin, "enable-insecure-login", false, "When enabled, Dashboard login view will also be shown when Dashboard is not served over HTTPS. Default: false.")
	serverCmd.PersistentFlags().StringVar(&argSystemBanner, "system-banner", "", "When non-empty displays message to Dashboard users. Accepts simple HTML tags. Default: ''.")
	serverCmd.PersistentFlags().StringVar(&argSystemBannerSeverity, "system-banner-severity", "INFO", "Severity of system banner. Should be one of 'INFO|WARNING|ERROR'. Default: 'INFO'.")

	serverCmd.PersistentFlags().StringVar(&argsHdfs, "hdfs", "http://localhost:50070", "")
	serverCmd.PersistentFlags().StringVar(&argsKuberUi, "kuber-ui", "http://localhost:4326", "")
	serverCmd.PersistentFlags().StringVar(&argsKuberRest, "kuber-rest", "http://localhost:9091", "")
	serverCmd.PersistentFlags().StringVar(&argsKuberWs, "kuber-ws", "ws://localhost:9091", "")

	serverCmd.PersistentFlags().StringVar(&argsGoogleApiKey, "google-api-key", "AIzaSyA4GOtTmfHmAL5t8jn0LBZ_SsInQukugAU", "")
	serverCmd.PersistentFlags().StringVar(&argsGoogleClientId, "google-client-id", "448379464054-clm37e3snnt3154cak4o5jqqmu4phs92.apps.googleusercontent.com", "")
	serverCmd.PersistentFlags().StringVar(&argsGoogleRedirect, "google-redirect", "http://localhost:9091/kuber/api/v1/google/redirect", "")
	serverCmd.PersistentFlags().StringVar(&argsGoogleSecret, "google-secret", "ZVxzNkOk98T2vEGbF5L-EQX3", "")
	serverCmd.PersistentFlags().StringVar(&argsGoogleScope, "google-scope", "profile email https://www.googleapis.com/auth/contacts.readonly https://www.googleapis.com/auth/user.emails.read https://www.googleapis.com/auth/gmail.insert", "")

	serverCmd.PersistentFlags().StringVar(&argsMicrosoftApplicationId, "microsoft-application-id", "86d75ba4-f7a0-4699-9c92-5c7a2bca194d", "")
	serverCmd.PersistentFlags().StringVar(&argsMicrosoftRedirect, "microsoft-redirect", "http://localhost:9091/kuber/api/v1/microsoft/callback", "")
	serverCmd.PersistentFlags().StringVar(&argsMicrosoftSecret, "microsoft-secret", "afMEQW2~?*jdyheSJU7715_", "")
	//	serverCmd.PersistentFlags().StringVar(&argsMicrosoftScope, "microsoft-scope", "User.ReadBasic.All+Contacts.Read+Mail.Send+Files.ReadWrite+Notes.ReadWrite", "")
	serverCmd.PersistentFlags().StringVar(&argsMicrosoftScope, "microsoft-scope", "User.ReadBasic.All", "")

	serverCmd.PersistentFlags().StringVar(&argsSpitfireRest, "spitfire-rest", "http://localhost:8091", "")
	serverCmd.PersistentFlags().StringVar(&argsSpitfireWs, "spitfire-ws", "ws://localhost:8091", "")

	serverCmd.PersistentFlags().StringVar(&argsTwitterConsumerKey, "twitter-consumer-key", "l8kmFysxMIpsSdmCrHa3qWL7d", "")
	serverCmd.PersistentFlags().StringVar(&argsTwitterConsumerSecret, "twitter-consumer-secret", "t4vRsq41orOzOnPMFYFMydqhgHMfe8NmstZndPUbySWsBTw0Mh", "")
	serverCmd.PersistentFlags().StringVar(&argsTwitterRedirect, "twitter-redirect", "", "")

	//	viper.BindPFlags(pflag.CommandLine)
	viper.BindPFlag("kuberui", serverCmd.PersistentFlags().Lookup("kuber-ui"))
	viper.BindPFlag("kuberrest", serverCmd.PersistentFlags().Lookup("kuber-rest"))
	viper.BindPFlag("kuberws", serverCmd.PersistentFlags().Lookup("kuber-ws"))
	viper.BindPFlag("googleclientid", serverCmd.PersistentFlags().Lookup("google-client-id"))
	viper.BindPFlag("googleredirect", serverCmd.PersistentFlags().Lookup("google-redirect"))
	viper.BindPFlag("googlesecret", serverCmd.PersistentFlags().Lookup("google-secret"))
	viper.BindPFlag("googlescope", serverCmd.PersistentFlags().Lookup("google-scope"))
	viper.BindPFlag("googleapikey", serverCmd.PersistentFlags().Lookup("google-api-key"))
	viper.BindPFlag("microsoftapplicationid", serverCmd.PersistentFlags().Lookup("microsoft-application-id"))
	viper.BindPFlag("microsoftredirect", serverCmd.PersistentFlags().Lookup("microsoft-redirect"))
	viper.BindPFlag("microsoftsecret", serverCmd.PersistentFlags().Lookup("microsoft-secret"))
	viper.BindPFlag("microsoftscope", serverCmd.PersistentFlags().Lookup("microsoft-scope"))
	viper.BindPFlag("spitfirerest", serverCmd.PersistentFlags().Lookup("spitfire-rest"))
	viper.BindPFlag("spitfirews", serverCmd.PersistentFlags().Lookup("spitfire-ws"))
	viper.BindPFlag("hdfs", serverCmd.PersistentFlags().Lookup("hdfs"))
	viper.BindPFlag("twitterconsumerkey", serverCmd.PersistentFlags().Lookup("twitter-consumer-key"))
	viper.BindPFlag("twitterconsumersecret", serverCmd.PersistentFlags().Lookup("twitter-consumer-secret"))
	viper.BindPFlag("twitterredirect", serverCmd.PersistentFlags().Lookup("twitter-redirect"))

}

func serve() {

	// Set logging output to standard console out.
	log.SetOutput(os.Stdout)

	//	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	//	pflag.Parse()

	flag.CommandLine.Parse(make([]string, 0)) // Init for glog calls in kubernetes packages.

	err := viper.Unmarshal(&config.KuberConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	log.Println("Kuber Config:", config.KuberConfig)

	log.Println("Starting WEB Server on port " + strconv.Itoa(argInsecurePort))

	go sanitize()
	serveWithK8s()

}

func serveAlone() {
	/*
		r := gorilla.Setup()
		log.Fatal(http.ListenAndServe(":"+*argPort,
			context.ClearHandler(
				handlers.CORS(gorilla.CredentialsOk(), gorilla.OriginsOk(), gorilla.HeadersOk(), gorilla.MethodsOk())(r))))
	*/
	wsContainer := restful.NewContainer()
	krest.SetupGoRestful(wsContainer)
	server := &http.Server{Addr: ":" + strconv.Itoa(argInsecurePort), Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

func serveWithK8s() {

	if argApiserverHost != "" {
		log.Printf("Using apiserver-host location: %s", argApiserverHost)
	}

	if argKubeConfigFile != "" {
		log.Printf("Using kubeconfig file: %s", argKubeConfigFile)
	}

	clientManager := client.NewClientManager(argKubeConfigFile, argApiserverHost)
	versionInfo, err := clientManager.InsecureClient().Discovery().ServerVersion()
	if err != nil {
		handleFatalInitError(err)
	}

	log.Printf("Successful initial request to the apiserver, version: %s", versionInfo.String())

	// Init auth manager.
	authManager := initAuthManager(clientManager, time.Duration(argTokenTTL))

	// Init settings manager.
	settingsManager := settings.NewSettingsManager(clientManager)

	// Init system banner manager.
	systemBannerManager := systembanner.NewSystemBannerManager(argSystemBanner, argSystemBannerSeverity)

	// Init integrations.
	integrationManager := integration.NewIntegrationManager(clientManager)
	integrationManager.Metric().ConfigureHeapster(argHeapsterHost).
		EnableWithRetry(integrationapi.HeapsterIntegrationID, time.Duration(argMetricClientCheckPeriod))

	apiHandler, err := handler.CreateHTTPAPIHandler(
		integrationManager,
		clientManager,
		authManager,
		argEnableInsecureLogin,
		settingsManager,
		systemBannerManager)
	if err != nil {
		handleFatalInitError(err)
	}

	krest.SetupGoRestful(apiHandler)

	if argAutoGenerateCertificates {
		log.Println("Auto-generating certificates")
		certCreator := ecdsa.NewECDSACreator(&argKeyFile, &argCertFile, elliptic.P256())
		certManager := cert.NewCertManager(certCreator, argDefaultCertDir)
		certManager.GenerateCertificates()
	}

	// Run a HTTP server that serves static public files from './public' and handles API calls.
	// TODO(bryk): Disable directory listing.
	//	http.Handle("/", handler.MakeGzipHandler(handler.CreateLocaleHandler()))

	http.Handle("/", handler.MakeGzipHandler(http.FileServer(http.Dir("./_ui/"))))

	http.Handle("/kuber/api/", apiHandler)
	// TODO(maciaszczykm): Move to /appConfig.json as it was discussed in #640.
	http.Handle("/kuber/api/appConfig.json", handler.AppHandler(handler.ConfigHandler))
	http.Handle("/kuber/api/sockjs/", handler.CreateAttachHandler("/kuber/api/sockjs"))
	http.Handle("/metrics", prometheus.Handler())

	// Listen for http or https
	if argCertFile != "" && argKeyFile != "" {
		certFilePath := argDefaultCertDir + string(os.PathSeparator) + argCertFile
		keyFilePath := argDefaultCertDir + string(os.PathSeparator) + argKeyFile
		log.Printf("Serving securely on HTTPS port: %d", argPort)
		secureAddr := fmt.Sprintf("%s:%d", argBindAddress, argPort)
		go func() { log.Fatal(http.ListenAndServeTLS(secureAddr, certFilePath, keyFilePath, nil)) }()
	} else {
		log.Printf("Serving insecurely on HTTP port: %d", argInsecurePort)
		addr := fmt.Sprintf("%s:%d", argInsecureBindAddress, argInsecurePort)
		go func() { log.Fatal(http.ListenAndServe(addr, nil)) }()
	}
	select {}
}

func initAuthManager(clientManager clientapi.ClientManager, tokenTTL time.Duration) authApi.AuthManager {
	insecureClient := clientManager.InsecureClient()

	// Init default encryption key synchronizer
	synchronizerManager := sync.NewSynchronizerManager(insecureClient)
	keySynchronizer := synchronizerManager.Secret(authApi.EncryptionKeyHolderNamespace, authApi.EncryptionKeyHolderName)

	// Register synchronizer. Overwatch will be responsible for restarting it in case of error.
	sync.Overwatch.RegisterSynchronizer(keySynchronizer, sync.AlwaysRestart)

	// Init encryption key holder and token manager
	keyHolder := jwe.NewRSAKeyHolder(keySynchronizer)
	tokenManager := jwe.NewJWETokenManager(keyHolder)
	if tokenTTL != authApi.DefaultTokenTTL {
		tokenManager.SetTokenTTL(tokenTTL)
	}

	// Set token manager for client manager.
	clientManager.SetTokenManager(tokenManager)
	authModes := authApi.ToAuthenticationModes(argAuthenticationMode)
	if len(authModes) == 0 {
		authModes.Add(authApi.Token)
	}

	return auth.NewAuthManager(clientManager, tokenManager, authModes)
}

/**
 * Handles fatal init error that prevents server from doing any work. Prints verbose error
 * message and quits the server.
 */
func handleFatalInitError(err error) {
	log.Fatalf("Error while initializing connection to Kubernetes apiserver. "+
		"This most likely means that the cluster is misconfigured (e.g., it has "+
		"invalid apiserver certificates or service accounts configuration) or the "+
		"--apiserver-host param points to a server that does not exist. Reason: %s\n"+
		"Refer to our FAQ and wiki pages for more information: "+
		"https://github.com/kubernetes/dashboard/wiki/FAQ", err)
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

	for true {
		if os.Getenv("KUBER_MANAGE_RESERVATIONS") == "true" {
			adjustNodeCapacity(region)
		}
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
