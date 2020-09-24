package config

import (
	"fmt"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

const (
	EnvApp   = "APP_ENV"
	AppLocal = "local"
)

const (
	ServicePort      = "service.port"
	PostgresUser     = "postgres.user"
	PostgresPassword = "postgres.password"
	PostgresDBname   = "postgres.dbname"
	PostgresPort     = "postgres.port"
	PostgresSslmode  = "postgres.sslmode"
	PostgresTimezone = "postgres.timezone"
)

func Init() {
	appEnv := os.Getenv(EnvApp)
	if appEnv == "" {
		appEnv = AppLocal
	}
	log.Printf("Reading config profile: %s", appEnv)

	if err := gotenv.Load(fmt.Sprintf("%s.env", appEnv)); err != nil {
		log.Fatalln("Failed to load env", err)
	}
}

func Get(name string) string {
	return os.Getenv(name)
}
