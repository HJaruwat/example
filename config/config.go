package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type (
	//Conf struct of
	Conf struct {
		App struct {
			Mysql   Mysql  `mapstructure:"mysql"`
			Amqp    Amqp   `mapstructure:"amqp"`
			GameURI string `mapstructure:"game-uri"`
		}
	}

	//Mysql struct of
	Mysql struct {
		Host     string `mapstructure:"host"`
		Database string `mapstructure:"database"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	}

	// Amqp struct of
	Amqp struct {
		Host        string `mapstructure:"host"`
		NameCash    string `mapstructure:"name-cash"`
		NameItem    string `mapstructure:"name-item"`
		NameMigrate string `mapstructure:"name-migrate"`
	}
)

// Setting is instance config setting
var Setting Conf

// Stage is environment program
var Stage string

func init() {
	setup()
	log.Printf(" [*] Env name : %s\n", Stage)
}

func setup() {
	Stage = os.Getenv("ENVIRONMENT")
	if Stage == "" {
		Stage = "development"
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName(fmt.Sprintf(".env.%s", Stage))

	viper.AddConfigPath(".")
	viper.AddConfigPath("/$GOPATH/src/cabal-api")
	viper.AddConfigPath("../")
	viper.AddConfigPath("/$HOME/GOPATH/src/cabal-api")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Setting)
	if err != nil {
		panic(err)
	}
}
