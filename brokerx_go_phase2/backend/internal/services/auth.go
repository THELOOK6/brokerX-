
package services

import (
	"errors"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

func HashPassword(pw string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pw), 12)
	if err != nil { return "", err }
	return string(h), nil
}

func CheckPassword(hash, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw)) == nil
}

func IssueJWT(sub, role string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" { return "", errors.New("missing JWT_SECRET") }
	claims := jwt.MapClaims{ "sub": sub, "role": role, "exp": time.Now().Add(2*time.Hour).Unix(), "iat": time.Now().Unix() }
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(secret))
}
