package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(err.Error())
	}
}

var (
	HttpPort      = os.Getenv("HTTP_PORT")
	DbHost        = os.Getenv("DB_HOST")
	DbPort        = os.Getenv("DB_PORT")
	DbUser        = os.Getenv("DB_USER")
	DbName        = os.Getenv("DB_NAME")
	DbPass        = os.Getenv("DB_PASSWORD")
	DbTz          = os.Getenv("DB_TZ")
	DbSslMode     = os.Getenv("DB_SSL_MODE")
	DbSslCertFile = os.Getenv("DB_SSL_CERT_FILE")
	DbSslKeyFile  = os.Getenv("DB_SSL_KEY_FILE")
)
