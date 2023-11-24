package internal

import (
	"Auth/internal/db"
	users "Auth/pb/proto"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

func Register(_ context.Context, r *users.UserRequest) (*users.UserResponse, error) {
	input := r.GetUserEntry()

	pass, err := HashPassword(input.Password)
	if err != nil {
		return &users.UserResponse{Result: "Error hashing the password"}, err
	}

	u := db.User{
		Id:       int(input.Id),
		Username: input.Username,
		Password: string(pass),
		Email:    input.Email,
		IsAdmin:  input.IsAdmin,
	}

	err = db.Create(u)
	if err != nil {
		return &users.UserResponse{Result: "Error creating the user"}, err
	}

	return &users.UserResponse{Result: "User created successfully"}, nil
}

func Login(_ context.Context, r *users.UserRequest) (*users.UserResponse, error) {
	input := r.GetUserEntry()

	uDb, err := db.GetByUsername(input.Username)
	if err != nil {
		return &users.UserResponse{Result: "Error getting the user in the database"}, err
	}

	if !CheckPasswordHash(input.Password, []byte(uDb.Password)) {
		return &users.UserResponse{Result: "Wrong password"}, errors.New("wrong password")
	}

	token, err := CreateJWTToken(uDb.Id, uDb.Username, uDb.IsAdmin)
	if err != nil {
		return &users.UserResponse{Result: "Error generating the token"}, err
	}

	return &users.UserResponse{Result: token}, nil
}

func VerifyJWT(_ context.Context, r *users.JWTRequest) (*users.JWTResponse, error) {
	tokenInput := r.GetToken()
	var userClaim UserClaim
	token, err := jwt.ParseWithClaims(tokenInput, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return &users.JWTResponse{Result: false}, err
	}
	if !token.Valid {
		return &users.JWTResponse{Result: false}, err
	}
	return &users.JWTResponse{Result: true}, nil
}
