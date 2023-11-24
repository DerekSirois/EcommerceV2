package internal

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

var secretKey = os.Getenv("SECRET_KEY")

type UserClaim struct {
	jwt.RegisteredClaims
	ID       int
	UserName string
	IsAdmin  bool
}

func CreateJWTToken(id int, name string, isAdmin bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24))},
		ID:               id,
		UserName:         name,
		IsAdmin:          isAdmin,
	})

	// Create the actual JWT token
	signedString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", fmt.Errorf("error creating signed string: %v", err)
	}

	return signedString, nil
}

func GetAuthenticatedUserId(r *http.Request) (int, error) {
	var jwtToken = r.Header["Token"][0]
	var userClaim UserClaim
	_, err := jwt.ParseWithClaims(jwtToken, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}
	return userClaim.ID, nil
}

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

func CheckPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
