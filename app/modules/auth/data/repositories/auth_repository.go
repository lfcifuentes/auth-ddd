package repositories

import (
	"errors"

	"github.com/lfcifuentes/auth-ddd/app/internal/adapters/pgsql"
)

type AuthRepository struct {
	DBAdapter *pgsql.DBAdapter
}

func NewAuthRepository(db *pgsql.DBAdapter) *AuthRepository {
	// Add your code here
	return &AuthRepository{
		DBAdapter: db,
	}
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthRepository) FindByEmail(email string) (*User, error) {

	rows := r.DBAdapter.DB.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email)
	var user User

	err := rows.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
