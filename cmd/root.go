package cmd

import (
	"fmt"
	"os"

	"io/ioutil"

	"github.com/mpppk/docker-env-wrapper/env"
	"github.com/spf13/cobra"
)

var checkImageFlag bool
var filterFlag string
var formatFlag string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "docker-env-wrapper",
	Short: "Generate Dockerfile or docker-compose.yml with host environment setting",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("image name must be specified")
			os.Exit(1)
		}

		env, err := env.New().Filter(filterFlag)
		if err != nil {
			panic(err)
		}

		out := fmt.Sprintln("FROM " + args[0])
		for k, v := range env {
			out += fmt.Sprintf("ENV %v %v\n", k, v)
		}
		ioutil.WriteFile("Dockerfile", []byte(out), 0777)
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

	RootCmd.Flags().BoolVarP(&checkImageFlag, "check-image", "c", false, "Check image is exist on Docker Hub")
	RootCmd.Flags().StringVarP(&filterFlag, "filter", "f", ".*", "Filter environment")
	RootCmd.Flags().StringVarP(&formatFlag, "format", "F", "dockerfile", "Specify output format")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}
