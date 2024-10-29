package configs

import (
	"os"

	"github.com/spf13/viper"
)

var (
	GRPC_PORT      string
	WEB_PORT       string
	DSN            string
	DB_USER        string
	DB_PASS        string
	DB_HOST        string
	DB_PORT        string
	DB_TABLE       string
	CONN_TRY_COUNT int
)

func init() {
	config_path := "."
	config_file_name := os.Getenv("CONFIG_FILE_NAME")
	config_file_type := os.Getenv("CONFIG_FILE_TYPE")
	viper.AddConfigPath(config_path)
	viper.SetConfigName(config_file_name)
	viper.SetConfigType(config_file_type)
	viper.ReadInConfig()

	// GRPC_PORT = viper.GetString("GRPC_PORT")
	// WEB_PORT = viper.GetString("WEB_PORT")
	// DB_USER = viper.GetString("DB_USER")
	// DB_PASS = viper.GetString("DB_PASS")
	// DB_HOST = viper.GetString("DB_HOST")
	// DB_PORT = viper.GetString("DB_PORT")
	// DB_TABLE = viper.GetString("DB_TABLE")
	// CONN_TRY_COUNT = viper.GetInt("CONN_TRY_COUNT")
	GRPC_PORT = "11012"
	WEB_PORT = "11011"
	DB_USER = "root"
	DB_PASS = "1234"
	DB_HOST = "127.0.0.1"
	DB_PORT = "11010"
	DB_TABLE = "myapp"
	CONN_TRY_COUNT = 24
}
