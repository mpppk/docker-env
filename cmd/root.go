package cmd

import (
	"fmt"
	"os"

	"path"

	"github.com/mpppk/docker-env-wrapper/env"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var checkImageFlag bool
var filterFlag string
var formatFlag string
var overwriteFlag bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "docker-env-wrapper",
	Short: "Generate Dockerfile or docker-compose.yml with host environment setting",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		env, err := env.New().Filter(filterFlag)
		if err != nil {
			panic(err)
		}
		fmt.Println(env)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	//RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.docker-env-wrapper.yaml)")
	RootCmd.Flags().BoolVarP(&checkImageFlag, "check-image", "c", false, "Check image is exist on Docker Hub")
	RootCmd.Flags().StringVarP(&filterFlag, "filter", "f", ".*", "Filter environment")
	RootCmd.Flags().StringVarP(&formatFlag, "format", "F", "dockerfile", "Specify output format")
	RootCmd.Flags().BoolVarP(&overwriteFlag, "overwrite", "o", false, "Overwrite output file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".docker-env-wrapper")                   // name of config file (without extension)
	viper.AddConfigPath(path.Join(os.Getenv("HOME"), ".config")) // adding home directory as first search path
	viper.AutomaticEnv()                                         // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
