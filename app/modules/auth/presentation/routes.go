package presentation

import (
	"github.com/lfcifuentes/auth-ddd/app/infrastructure"
	"github.com/lfcifuentes/auth-ddd/app/modules/auth/domain/services"
)

func RegisterAuthRoutes(a *infrastructure.Application) {
	authServices := services.NewAuthServices(a.DbAdapter, a.Validator)

	authGroup := a.Router.Group("/auth")

	authGroup.POST("/login", authServices.Login)
}
