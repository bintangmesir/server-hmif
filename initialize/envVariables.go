package initialize

import (
	"os"

	"github.com/joho/godotenv"
)

var (
    ENV_PORT string
    ENV_CLIENT_URI string    
    ENV_POTOCOL_HTTPS string

    ENV_ACCESS_TOKEN string
    ENV_REFRESH_TOKEN string

    ENV_DB_URI string
    ENV_DB_PORT string
    ENV_DB_NAME string
    ENV_DB_USER string
    ENV_DB_PASSWORD string

    ENV_DIR_HIMPUNAN_FILES string
    ENV_DIR_ADMIN_FILES string
    ENV_DIR_BUKU_FILES string
    ENV_DIR_ARTIKEL_FILES string
    ENV_DIR_PENGURUS_FILES string
    ENV_DIR_COMMENT_FILES string
    ENV_DIR_ARTIKEL_CONTENT_FILES string
)

func EnvVariables (){
    godotenv.Load()

    ENV_PORT = os.Getenv("ENV_PORT")
    ENV_CLIENT_URI = os.Getenv("ENV_CLIENT_URI")    
    ENV_POTOCOL_HTTPS = os.Getenv("ENV_POTOCOL_HTTPS")

    ENV_ACCESS_TOKEN = os.Getenv("ENV_ACCESS_TOKEN")
    ENV_REFRESH_TOKEN = os.Getenv("ENV_REFRESH_TOKEN")

    ENV_DB_URI = os.Getenv("ENV_DB_URI")
    ENV_DB_PORT = os.Getenv("ENV_DB_PORT")
    ENV_DB_NAME = os.Getenv("ENV_DB_NAME")
    ENV_DB_USER = os.Getenv("ENV_DB_USER")
    ENV_DB_PASSWORD = os.Getenv("ENV_DB_PASSWORD")

    ENV_DIR_HIMPUNAN_FILES = os.Getenv("ENV_DIR_HIMPUNAN_FILES")
    ENV_DIR_ADMIN_FILES = os.Getenv("ENV_DIR_ADMIN_FILES")
    ENV_DIR_BUKU_FILES = os.Getenv("ENV_DIR_BUKU_FILES")
    ENV_DIR_PENGURUS_FILES = os.Getenv("ENV_DIR_PENGURUS_FILES")
    ENV_DIR_ARTIKEL_FILES = os.Getenv("ENV_DIR_ARTIKEL_FILES")
    ENV_DIR_COMMENT_FILES = os.Getenv("ENV_DIR_COMMENT_FILES")
    ENV_DIR_ARTIKEL_CONTENT_FILES = os.Getenv("ENV_DIR_ARTIKEL_CONTENT_FILES")
}

