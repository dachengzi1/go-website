package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type RedisConfig struct {
	Ip   string `toml:"ip"`
	Port int    `toml:"port"`
	Db   int    `toml:"db"`
}

type MysqlConfig struct {
	Ip       string `toml:"ip"`
	Port     int    `toml:"port"`
	Database string `toml:"database"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type Config struct {
	Env   string
	Redis RedisConfig
	Mysql MysqlConfig
}

var Conf Config

func init() {
	env := os.Getenv("GO_ENV")
	fmt.Printf("go env is :%v\n", env)
	viper.SetConfigName(env)
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}
	viper.Unmarshal(&Conf)
	fmt.Println("config:", Conf)
}
