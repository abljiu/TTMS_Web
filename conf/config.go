package conf

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Service struct {
		AppMode  string `yaml:"AppMode"`
		HttpPort int    `yaml:"HttpPort"`
	} `yaml:"service"`
	Mysql struct {
		DB         string `yaml:"DB"`
		DBHost     string `yaml:"DBHost"`
		DbPort     string `yaml:"DbPort"`
		DbPassword string `yaml:"DbPassword"`
		DbName     string `yaml:"DbName"`
	} `yaml:"mysql"`
	Redis struct {
		RedisDb     string `yaml:"RedisDb"`
		RedisAddr   string `yaml:"RedisAddr"`
		RedisPw     string `yaml:"RedisPw"`
		RedisDbName int    `yaml:"RedisDbName"`
	} `yaml:"redis"`
	Email struct {
		ValidEmail string `yaml:"ValidEmail"`
		SmtpHost   string `yaml:"SmtpHost"`
		SmtpEmail  string `yaml:"SmtpEmail"`
		SmtpPass   string `yaml:"SmtpPass"`
	} `yaml:"email"`
	Path struct {
		Host        string `yaml:"Host"`
		ProductPath string `yaml:"ProductPath"`
		AvatarPath  string `yaml:"AvatarPath"`
	} `yaml:"path"`
}

func Init() {
	var config Config

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

}
