package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	HttpPort string
	APP      string
	ENV      string
	DB       Database
}

type Database struct {
	DB_DRIVER string `structure:"DB_DRIVER"`
	DB_USER   string `structure:"DB_USER"`
	DB_PASS   string `structure:"DB_PASS"`
	DB_HOST   string `structure:"DB_HOST"`
	DB_NAME   string `structure:"DB_NAME"`
	DB_PORT   string `structure:"DB_PORT"`
}

var (
	Values Config
	DB     Database
)

func init() {
	// Config file extension
	viper.SetConfigType("env")
	// Include env vars (Env vars precedence is bigger than config files)
	viper.AutomaticEnv()
	// Config file name
	viper.SetConfigName("conf")
	viper.SetConfigFile("./.env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&DB)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	err = viper.Unmarshal(&Values)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	Values.DB = DB
}

func (d Database) ConnectionString() string {
	myConnString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	if os.Getenv("DB_HOST") == "" || os.Getenv("DB_USER") == "" || os.Getenv("DB_PASS") == "" || os.Getenv("DB_NAME") == "" {
		myConnString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", d.DB_HOST, d.DB_USER, d.DB_PASS, d.DB_PORT, d.DB_NAME)
		if d.DB_HOST == "" || d.DB_USER == "" || d.DB_PASS == "" || d.DB_NAME == "" {
			panic(errors.New("the connection string is empty or not found"))
		}
	}
	return myConnString
}
