package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	GRPC_PORT      string
	WEB_PORT       string
	DSN            string
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
	log.Printf("Setting config using file '%v/%v.%v", config_path, config_file_name, config_file_type)

	GRPC_PORT = viper.GetString("GRPC_PORT")
	GRPC_PORT = "54321"
	// WEB_PORT = viper.GetString("WEB_PORT")
	WEB_PORT = "12345"
	// DSN = viper.GetString("DATA_SOURCE_NAME")
	DSN = "root:1234@tcp(127.0.0.1:11010)/myapp"
	CONN_TRY_COUNT = viper.GetInt("CONN_TRY_COUNT")
}
