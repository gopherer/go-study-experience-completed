package tools

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName  string         `json:"app_name"`
	AppMode  string         `json:"app_mode"`
	AppHost  string         `json:"app_host"`
	AppPort  string         `json:"app_port"`
	Sms      SmsConfig      `json:"sms"`
	Database DatabaseConfig `json:"database"`
}

type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	RegionID     string `json:"region_id"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	CharSet  string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

var cfg *Config

func GetConfig() *Config {
	return cfg
}
func ParseConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err.Error())
		}
	}()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	} else {
		return cfg, nil
	}
}
