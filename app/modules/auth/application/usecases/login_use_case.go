package usecases

import (
	"errors"

	"github.com/lfcifuentes/auth-ddd/app/modules/auth/data/repositories"
	"github.com/lfcifuentes/auth-ddd/app/modules/auth/data/valueobjects"
	"github.com/lfcifuentes/auth-ddd/app/services"
)

type LoginUseCase struct {
	repo           *repositories.AuthRepository
	passwordHasher services.PasswordHasher
	jwt            services.ApiJWT
}

func NewLoginUseCase(repo *repositories.AuthRepository) *LoginUseCase {
	return &LoginUseCase{
		repo:           repo,
		passwordHasher: *services.NewPasswordHasher(),
		jwt:            *services.NewApiJWT(),
	}
}

func (uc *LoginUseCase) Login(email, password string) (*valueobjects.LoginResponse, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !uc.passwordHasher.Check(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	token, expiration, err := uc.jwt.GenerateJWT(user.ID)

	if err != nil {
		return nil, errors.New("error generating token")
	}

	err = uc.repo.SaveToken(user.ID, token, expiration)

	if err != nil {
		return nil, errors.New("error saving token")
	}

	// Return user data or token as needed
	return &valueobjects.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   expiration,
	}, nil
}
