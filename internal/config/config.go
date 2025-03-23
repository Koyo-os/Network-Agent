package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Notify struct {
	Send   bool   `yaml:"send"`
	ChatId int64  `yaml:"chat_id"`
	Token  string `yaml:"token"`
}

type Build struct {
	InputPoint  string `yaml:"input"`
	OutputPoint string `yaml:"output"`
}

type Config struct {
	Port        int `yaml:"port"`
	Host        string `yaml:"host"`
	TempDirName string `yaml:"temp_dir_name"`
	BuildCfg    Build  `yaml:"build"`
	NotifyCfg   Notify `yaml:"notify"`
}

func Init() (*Config, error) {
	file, err := os.Open("config.yml")
	if err != nil{
		return nil, fmt.Errorf("cant open config file: %v",err)
	}

	var cfg Config
	if err = yaml.NewDecoder(file).Decode(&cfg);err != nil{
		return nil, fmt.Errorf("cant decode config: %v", err)
	}

	return &cfg, nil
}