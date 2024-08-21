package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfcifuentes/auth-ddd/app/infrastructure"
	"github.com/lfcifuentes/auth-ddd/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/auth-ddd/docs"
	"github.com/spf13/viper"
)

func NewRouter(db *pgsql.DBAdapter) *gin.Engine {
	// Add swagger info
	docs.SwaggerInfo.Host = viper.GetString("APP_HOST")
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Version = viper.GetString("APP_VERSION")

	// create the application
	app := infrastructure.NewApplication(
		db,
	)

	// Add your code here

	return app.Router

}
