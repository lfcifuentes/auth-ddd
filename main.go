package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/lfcifuentes/auth-ddd/app/cmd"
	"github.com/spf13/viper"
)

//	@title			Simple Auth API documentation
//	@version		1.0
//	@description	This is a simple auth API documentation.
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	Luis Cifuentes
//	@contact.url	https://lfcifuentes.netlify.app
//	@contact.email	lfcifuentes28@gmail.com
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host
//	@Accept		json
//	@Produce	json
//	@BasePath	/

func main() {
	// Create a new logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	// Show the config file path
	slog.Info("Starting application", "event", "app_starting")

	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Configurar Viper para leer variables de entorno
	viper.AutomaticEnv()
	// Show the config file path
	slog.Info("Using config file", "event", "app_env_file_load", "file", ".env")

	cmd.Execute()
}
