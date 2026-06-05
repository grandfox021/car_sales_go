package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Logger   LoggerConfig

}

type ServerConfig struct {
	InternalPort string
	ExternalPort string
	RunMode      string
	Domain       string
	Port	     string
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
	Logger   string
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	Db                 string
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
}




func GetConfig() *Config{

	cfg_path := get_config_path(os.Getenv("APP_ENV"))
	v,err := ReadConfig(cfg_path , "yaml")
	if err!=nil {

		log.Fatal(err)
	}

	cfg,err := ParsConfig(v)
	if err!=nil {

		log.Fatal(err)
	}
	return cfg

}


func ParsConfig(v *viper.Viper ) (*Config ,error){

	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {

		log.Printf("unable to parse config err: %s",err)
		return nil , err

	}

	return &cfg , nil

}



func ReadConfig(filename string ,filetype string)(*viper.Viper , error){

	v := viper.New()
	v.SetConfigType(filetype)
	v.SetConfigFile(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()

	if err != nil{
		if _, ok := err.(viper.ConfigFileNotFoundError); ok{
			return nil , errors.New("configfile not found !")
		}
		return  nil,err
	}
	return v , nil
}



func get_config_path(env string ) string{

	if env == "docker"{
		return "config/config-docker"
	}else if env == "production"{
		return "config/config-production"
	}else {
		return  "../config/config-development.yaml"
	}
}