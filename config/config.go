package config

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"

	"go-ent-gqlgen/pkg/util/environment"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Database string
}

type ServerConfig struct {
	Address string
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

// C is config variable
var C Config

// ReadConfigOption is a config option
type ReadConfigOption struct {
	AppEnv string
}

// ReadConfig configures config file
func ReadConfig(option ReadConfigOption) error {
	config := &C

	if environment.IsDev() {
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config")
	} else if environment.IsTest() || (option.AppEnv == environment.Test) {
		fmt.Println(rootDir())
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config.test")
	} else if environment.IsE2E() || (option.AppEnv == environment.E2E) {
		fmt.Println(rootDir())
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config.e2e")
	} else {
		// production configuration here
		fmt.Println("production configuration here")
	}

	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
