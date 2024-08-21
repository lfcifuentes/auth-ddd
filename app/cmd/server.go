package cmd

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/lfcifuentes/auth-ddd/app/infrastructure/http/router"
	"github.com/lfcifuentes/auth-ddd/app/internal/adapters/pgsql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(runServer)
}

var runServer = &cobra.Command{
	Use:   "server",
	Short: "Run app server",
	Run:   runServerHandler,
}

func runServerHandler(_ *cobra.Command, _ []string) {
	viper.SetDefault("HTTP_PORT", "8001")
	viper.SetDefault("APP_VERSION", "1.0.0")
	// Create a new mongodb adapter
	slog.Info("Creating database adapter", "event", "database_adapter_creation")
	db, err := pgsql.NewDBAdapter()
	if err != nil {
		slog.Error(err.Error(), "event", "database_connection_failed")
		return
	}
	err = db.Ping()
	if err != nil {
		slog.Error(err.Error(), "event", "database_pin_connection_failed")
		return
	}

	slog.Info("Creating server", "event", "server_creation")
	// Configura el enrutador
	routes := router.NewRouter(db)

	slog.Info("Server created", "event", "server_created")
	port := viper.GetString("HTTP_PORT")
	// Inicia el servidor web
	http.ListenAndServe(
		fmt.Sprintf(":%v", port),
		routes,
	)
}
