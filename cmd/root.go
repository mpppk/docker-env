package cmd

import (
	"fmt"
	"os"

	"io/ioutil"

	"strings"

	"strconv"

	"github.com/mpppk/docker-env/env"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var queryFlag string
var formatFlag string

const FORMAT_fLAG_DOCKER_FILE = "dockerfile"
const FORMAT_fLAG_DOCKER_COMPOSE = "compose"
const DOCKER_FILE_NAME = "Dockerfile"
const DOCKER_COMPOSE_FILE_NAME = "docker-compose.yml"

var RootCmd = &cobra.Command{
	Use:   "docker-env IMAGE_NAME",
	Short: "Generate Dockerfile or docker-compose.yml with host environment variables",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Image name must be specified")
			os.Exit(1)
		}

		env, err := env.New().Filter(queryFlag)
		if err != nil {
			panic(err)
		}

		switch formatFlag {
		case FORMAT_fLAG_DOCKER_FILE:
			var out string
			for k, v := range env {
				out += fmt.Sprintf("ENV %v %v\n", k, v)
			}

			for i, imageName := range args {
				content := fmt.Sprintf("FROM %v\n%v", imageName, out)
				fmt.Println(content)
				suffix := strconv.Itoa(i)
				if i == 0 {
					suffix = ""
				}
				ioutil.WriteFile(DOCKER_FILE_NAME+suffix, []byte(content), 0777)
			}

		case FORMAT_fLAG_DOCKER_COMPOSE:
			services := map[string]interface{}{}
			for _, imageName := range args {
				containerKey := strings.Replace(imageName, ":", "", -1)
				services[containerKey] = map[string]interface{}{
					"image":       imageName,
					"environment": env,
				}
			}

			d := map[string]interface{}{
				"version":  "3",
				"services": services,
			}
			o, err := yaml.Marshal(d)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(o))
			ioutil.WriteFile(DOCKER_COMPOSE_FILE_NAME, o, 0777)
		}
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

	RootCmd.Flags().StringVarP(&queryFlag, "query", "q", ".*", "Filter host environment variables with regular expressions.")
	RootCmd.Flags().StringVarP(&formatFlag, "format", "f", FORMAT_fLAG_DOCKER_FILE, "Specify output format. [dockerfile|compose]")
}

func initConfig() {
}
