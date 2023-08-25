package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/muhwyndhamhp/marknotes/utils/errs"
)

const (
	ENV_FILE = ".env"

	APP_PORT              = "APP_PORT"
	DB_HOST               = "DB_HOST"
	DB_PORT               = "DB_PORT"
	DB_USER               = "DB_USER"
	DB_NAME               = "DB_NAME"
	DB_PASSWORD           = "DB_PASSWORD"
	JWT_SECRET            = "JWT_SECRET"
	OAUTH_AUTHORIZE_URL   = "OAUTH_AUTHORIZE_URL"
	OAUTH_ACCESSTOKEN_URL = "OAUTH_ACCESSTOKEN_URL"
	OAUTH_CLIENTID        = "OAUTH_CLIENTID"
	OAUTH_SECRET          = "OAUTH_SECRET"
	OAUTH_URL             = "OAUTH_URL"
)

func init() {
	if err := godotenv.Load(ENV_FILE); err != nil {
		log.Fatal(errs.Wrap(err))
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
