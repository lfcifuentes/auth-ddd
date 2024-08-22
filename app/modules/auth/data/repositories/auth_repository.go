package repositories

import (
	"errors"
	"time"

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

// FindByEmail retrieves a user by email.
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

// SaveToken saves a token in the database.
func (r *AuthRepository) SaveToken(id int, token string, expirationTime time.Time) error {
	_, err := r.DBAdapter.DB.Exec("INSERT INTO tokens (user_id, token, expires_at, status) VALUES ($1, $2, $3, $4)", id, token, expirationTime, true)
	if err != nil {
		return errors.New("error saving token")
	}

	return nil
}

// ValidateToken checks if a token is valid.
func (r *AuthRepository) ValidateToken(token string) (bool, error) {
	rows := r.DBAdapter.DB.QueryRow("SELECT status FROM tokens WHERE token = $1", token)
	var status bool

	err := rows.Scan(&status)
	if err != nil {
		return false, errors.New("token not found")
	}

	return status, nil
}

// InvalidateToken invalidates a token.
func (r *AuthRepository) InvalidateToken(token string) error {
	_, err := r.DBAdapter.DB.Exec("UPDATE tokens SET status = false WHERE token = $1", token)
	if err != nil {
		return errors.New("error invalidating token")
	}

	return nil
}
