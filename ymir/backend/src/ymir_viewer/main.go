package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/IndustryEssentials/ymir-viewer/common/constants"
	"github.com/IndustryEssentials/ymir-viewer/services"
)

func InitViperConfig() constants.Config {
	viper.BindEnv("YmirSandbox", "BACKEND_SANDBOX_ROOT")
	viper.BindEnv("ViewerPort", "VIEWER_HOST_PORT")
	viper.BindEnv("MongodbUri", "MONGODB_URI")

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config constants.Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	config.ViewerUri = fmt.Sprintf("%s:%d", config.ViewerHost, config.ViewerPort)
	log.Printf("config: %+v\n", config)
	return config
}

func main() {
	viewerServer := services.NewViewerServer(InitViperConfig())
	defer viewerServer.Clear()

	viewerServer.Start()
}