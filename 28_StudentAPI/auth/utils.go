package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Secret key used for signing JWT tokens. Keep this private in real apps.
var JwtKey = []byte("abcdgduchjdhni")

type Claims struct {
	UserId               uint
	Email                string
	jwt.RegisteredClaims // includes Expiry, IssuedAt etc.
}

// Hashes the plain password using bcrypt before saving it
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}


// Compares hashed password (from DB) and plain password (from login form)
func CheckPasswordHash(hashedPassword, plainPassword string) bool  {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil // if no error then return true
}

// Creates a JWT token with userID & email, expires in 24h
func GenerateJWT(userID uint, email string) (string, error) {
	claims := Claims{
		UserId: userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour *24)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString(JwtKey) // sign using secret key
}