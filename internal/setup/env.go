package setup

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	Mode       string `mapstructure:"APP_ENV"`
	Address    string `mapstructure:"SERVER_ADDRESS"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
	DBUser     string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBPort     string `mapstructure:"POSTGRES_PORT"`
	DBName     string `mapstructure:"POSTGRES_NAME"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Can't find the config .env file: %s", err)
	}

	err = viper.Unmarshal(&env)

	if err != nil {
		log.Fatalf("Error while reading env file: %s", err)
	}

	if env.Mode == "development" {
		log.Println("App running in development mode, take it in view.")
	}

	return &env
}
