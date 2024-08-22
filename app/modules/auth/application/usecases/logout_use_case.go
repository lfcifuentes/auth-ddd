package usecases

import (
	"errors"

	"github.com/lfcifuentes/auth-ddd/app/modules/auth/data/repositories"
	"github.com/lfcifuentes/auth-ddd/app/services"
)

type LogoutUseCase struct {
	repo *repositories.AuthRepository
	jwt  services.ApiJWT
}

func NewLogoutUseCase(repo *repositories.AuthRepository) *LogoutUseCase {
	// Add your code here
	return &LogoutUseCase{
		repo: repo,
		jwt:  *services.NewApiJWT(),
	}
}

func (uc *LogoutUseCase) Logout(token string) error {
	_, token, err := uc.jwt.ValidateJWT(token)
	if err != nil {
		return err
	}

	tokenStatus, err := uc.repo.ValidateToken(token)
	if err != nil {
		return err
	}

	if !tokenStatus {
		return errors.New("invalid token")
	}

	err = uc.repo.InvalidateToken(token)
	if err != nil {
		return err
	}
	// Add your code here
	return nil
}
