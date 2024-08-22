package usecases

import (
	"errors"

	"github.com/lfcifuentes/auth-ddd/app/modules/auth/data/repositories"
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

func (uc *LoginUseCase) Login(email, password string) (interface{}, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !uc.passwordHasher.Check(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	token, err := uc.jwt.GenerateJWT(user.ID)

	// Return user data or token as needed
	return map[string]interface{}{
		"token":   token,
		"user_id": user.ID,
		"email":   user.Email,
	}, nil
}
