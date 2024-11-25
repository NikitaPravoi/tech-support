package services

import (
	"context"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"tech-support/internal/database"
	"time"
)

// HashPassword generates a bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a bcrypt hashed password with its possible plaintext equivalent
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func CheckTokenValid(ctx context.Context, token string) *database.Session {
	session, err := database.Repository.GetSessionByToken(ctx, token)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil
		}
		return nil
	}

	if session.ExpiresAt.Time.Before(time.Now()) {
		return nil
	}

	return &session
}
