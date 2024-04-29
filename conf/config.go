package conf

import (
	"TTMS_Web/dao"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

var Config_ Config

type Config struct {
	Service struct {
		AppMode  string `yaml:"AppMode"`
		HttpPort string `yaml:"HttpPort"`
	} `yaml:"service"`
	Mysql struct {
		DB         string `yaml:"DB"`
		DbHost     string `yaml:"DbHost"`
		DbPort     string `yaml:"DbPort"`
		DbUser     string `yaml:"DbUser"`
		DbPassword string `yaml:"DbPassword"`
		DbName     string `yaml:"DbName"`
	} `yaml:"mysql"`
	Redis struct {
		RedisDb     string `yaml:"RedisDb"`
		RedisAddr   string `yaml:"RedisAddr"`
		RedisPw     string `yaml:"RedisPw"`
		RedisDbName string `yaml:"RedisDbName"`
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
	yamlFile, err := os.ReadFile("./conf/config.yaml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &Config_)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	pathRead := strings.Join([]string{Config_.Mysql.DbUser, ":", Config_.Mysql.DbPassword, "@tcp(", Config_.Mysql.DbHost, ":", Config_.Mysql.DbPort, ")/", Config_.Mysql.DbName, "?charset=utf8mb4&parseTime=true"}, "")
	pathWrite := strings.Join([]string{Config_.Mysql.DbUser, ":", Config_.Mysql.DbPassword, "@tcp(", Config_.Mysql.DbHost, ":", Config_.Mysql.DbPort, ")/", Config_.Mysql.DbName, "?charset=utf8mb4&parseTime=true"}, "")
	dao.Database(pathRead, pathWrite)
}
