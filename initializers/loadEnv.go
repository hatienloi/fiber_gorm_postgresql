package initializers

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

type ConfigStruct struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`
	TokenExpiresIn string `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenSecret    string `mapstructure:"TOKEN_SECRET"`
}

var Config ConfigStruct

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error loading .env file")
	}
	myEnv, _ := godotenv.Read()
	err = mapstructure.Decode(myEnv, &Config)
	if err != nil {
		panic("not enough variable environment")
	}
	fmt.Println(Config)
}
