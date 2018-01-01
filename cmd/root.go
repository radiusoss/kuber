package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/kris-nova/kubicorn/cutil/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "kuber",
	Short: "kuber, Easy Kubernetes",
	Long: `
Kuber allows you to create and manage a Kubernetes cluster for Big Data Science needs.`,
	// Uncomment the following line if your bare application has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here, will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kuber.yaml)")

	RootCmd.PersistentFlags().IntVarP(&logger.Level, "verbose", "v", 3, "Log level")
	RootCmd.PersistentFlags().BoolVarP(&logger.Color, "color", "C", true, "Toggle colorized logs")
	RootCmd.PersistentFlags().BoolVarP(&logger.Fabulous, "fab", "f", false, "Toggle colorized logs")

	// Cobra also supports local flags, which will only run when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// add commands

	// add commands
	addCommands()
}

func addCommands() {
	RootCmd.AddCommand(CreateCmd())
	RootCmd.AddCommand(ApplyCmd())
	RootCmd.AddCommand(DeleteCmd())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	/*
		if cfgFile != "" {
			// Use config file from the flag.
			viper.SetConfigFile(cfgFile)
		} else {
			// Find home directory.
			home, err := homedir.Dir()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Search config in home directory with name ".cobra-example" (without extension).
			viper.AddConfigPath(home)
			viper.SetConfigName(".kuber")
		}

		viper.AutomaticEnv() // read in environment variables that match

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	*/
	viper.AddConfigPath("/etc/datalayer/kuber")
	viper.AddConfigPath("$HOME/.datalayer/kuber")
	viper.AddConfigPath("./kuber")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	viper.SetConfigName("kuber-conf")
	viper.SetConfigType("yml")

	// Find and read the config file.
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Error reading config file", err)
	}
	// Confirm which config file is used.
	log.Printf("Using config file: %s\n", viper.ConfigFileUsed())

	viper.SetEnvPrefix("kuber")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

}

func flagApplyAnnotations(cmd *cobra.Command, flag, completion string) {
	if cmd.Flag(flag) != nil {
		if cmd.Flag(flag).Annotations == nil {
			cmd.Flag(flag).Annotations = map[string][]string{}
		}
		cmd.Flag(flag).Annotations[cobra.BashCompCustom] = append(
			cmd.Flag(flag).Annotations[cobra.BashCompCustom],
			completion,
		)
	}
}
