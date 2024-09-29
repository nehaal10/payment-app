package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server Server `json:"server"`
	DB     DB     `json:"db"`
}

type Server struct {
	Address        string `json:"address"`
	SessionTimeout string `json:"session_timeout"`
}

type DB struct {
	Type string `json:"type"`
}

func LoadConfig(env string) (config *Config, err error) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	properPath := fmt.Sprintf("%s/config/config.%s.json", dir, env)

	byteData, err := os.ReadFile(properPath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = json.Unmarshal(byteData, &config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return config, nil
}
