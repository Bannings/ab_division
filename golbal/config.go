package golbal

import (
	"encoding/json"
	"io/ioutil"
)

var (
	conf *Config
)

type Config struct {
	ABMysql DBConfig    `json:"ab_mysql"`
	Redis   RedisConfig `json:"redis_cli"`
	Port    string      `json:"port"`
}

type DBConfig struct {
	Schema   string `json:"schema"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
	PoolSize int    `json:"pool_size"`
}

func LoadConfig(jsonFile string) (*Config, error) {
	file, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	tmp := new(Config)
	if err = json.Unmarshal(file, tmp); err != nil {
		return nil, err
	}

	conf = tmp
	return conf, nil
}

func GetConfig() *Config {
	return conf
}
