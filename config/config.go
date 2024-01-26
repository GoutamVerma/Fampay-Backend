package config

import (
	"fmt"
	"log"

	viper "github.com/spf13/viper"
)

func ReadUserName() string {
	viper.SetConfigFile("backend-server.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetString("username")
}
func ReadPassWord() string {
	viper.SetConfigFile("backend-server.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetString("password")
}

func ReadHostName() string {
	viper.SetConfigFile("backend-server.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetString("hostname")
}

func ReadDatabaseName() string {
	viper.SetConfigFile("backend-server.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetString("databaseName")
}

func ReadPort() int {
	viper.SetConfigFile("backend-server.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetInt("port")
}

func FetchInternval() int {
	viper.SetConfigFile("backend-server.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetInt("FetchInterval")
}

func ReadYouTubeAPIKeys() []string {
	var apiKeys []string

	totalAPIKeys := viper.GetInt("TotalAPIKeys")
	for i := 1; i <= totalAPIKeys; i++ {
		key := viper.GetString(fmt.Sprintf("youtubeAPI%d", i))
		if key != "" {
			apiKeys = append(apiKeys, key)
		}
	}

	if len(apiKeys) == 0 {
		log.Fatal("No valid YouTube API keys found in the configuration.")
	}

	return apiKeys
}
