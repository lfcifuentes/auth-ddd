package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfcifuentes/auth-ddd/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/auth-ddd/app/internal/adapters/validator"
	"github.com/lfcifuentes/auth-ddd/app/modules/auth/application/usecases"
	"github.com/lfcifuentes/auth-ddd/app/modules/auth/data/repositories"
	"github.com/lfcifuentes/auth-ddd/app/modules/auth/data/valueobjects"
)

type AuthServices struct {
	Repository    *repositories.AuthRepository
	Validator     *validator.ValidatorAdapter
	LoginUseCase  *usecases.LoginUseCase
	LogoutUseCase *usecases.LogoutUseCase
}

func NewAuthServices(db *pgsql.DBAdapter, validator *validator.ValidatorAdapter) *AuthServices {
	repo := repositories.NewAuthRepository(db)
	// Add your code here
	return &AuthServices{
		Validator:     validator,
		LoginUseCase:  usecases.NewLoginUseCase(repo),
		LogoutUseCase: usecases.NewLogoutUseCase(repo),
	}
}

// Login Loguear un usuario usado nuestra base de datos
//
//	@Summary		Loguear un usuario usado nuestra base de datos
//	@Description	Loguear un usuario usado nuestra base de datos
//	@Tags			Auth
//	@Param			userCredentials	body	valueobjects.LoginParams	true	"User Credentials"
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	gin.H{data=nil}
//	@Failure		400	{object}	gin.H{data=nil}
//	@Failure		401	{object}	gin.H{data=nil}
//	@Failure		404	{object}	gin.H{data=nil}
//	@Failure		500	{object}	gin.H{data=nil}
//	@Router			/auth/login [post]
func (s *AuthServices) Login(c *gin.Context) {
	var request valueobjects.LoginParams
	_ = c.ShouldBindJSON(&request)

	errStruct := s.Validator.ValidateStruct(request)
	if errStruct != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": errStruct},
		)
		return
	}

	data, err := s.LoginUseCase.Login(
		request.Email,
		request.Password,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"data": data},
	)
}

// Logout Desloguear un usuario usado nuestra base de datos
//
//	@Summary		Desloguear un usuario usado nuestra base de datos
//	@Description	Desloguear un usuario usado nuestra base de datos
//	@Tags			Auth
//	@Param			token	header	string	true	"Token"
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	gin.H{data=nil}
//	@Failure		400	{object}	gin.H{data=nil}
//	@Failure		401	{object}	gin.H{data=nil}
//	@Failure		404	{object}	gin.H{data=nil}
//	@Failure		500	{object}	gin.H{data=nil}
//	@Router			/auth/logout [post]
func (s *AuthServices) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")

	err := s.LogoutUseCase.Logout(token)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"data": "User logged out successfully"},
	)
}
