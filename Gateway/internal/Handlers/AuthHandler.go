package Handlers

import (
	"context"
	"errors"
	users "gateway/pb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
)

type User struct {
	Id       int
	Username string
	Password string
	Email    string
	IsAdmin  bool
}

func Register(w http.ResponseWriter, r *http.Request) {
	var requestPayload User

	err := readJson(r, &requestPayload)
	if err != nil {
		errorJson(w, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	conn, err := grpc.Dial("auth:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		errorJson(w, http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	defer conn.Close()

	c := users.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Register(ctx, &users.UserRequest{
		UserEntry: &users.User{
			Id:       int32(requestPayload.Id),
			Username: requestPayload.Username,
			Password: requestPayload.Password,
			Email:    requestPayload.Email,
			IsAdmin:  false,
		},
	})
	if err != nil {
		errorJson(w, http.StatusInternalServerError, err)
		log.Println(err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: res.GetResult(),
	}

	_ = sendJson(w, http.StatusOK, payload)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var requestPayload User
	err := readJson(r, &requestPayload)
	if err != nil {
		errorJson(w, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	conn, err := grpc.Dial("auth:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		errorJson(w, http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	defer conn.Close()

	c := users.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Login(ctx, &users.UserRequest{
		UserEntry: &users.User{
			Username: requestPayload.Username,
			Password: requestPayload.Password,
		},
	})
	if err != nil {
		errorJson(w, http.StatusInternalServerError, err)
		log.Println(err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: res.GetResult(),
	}

	_ = sendJson(w, http.StatusOK, payload)
}

func VerifyJWT(endpointHandler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			conn, err := grpc.Dial("auth:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
			if err != nil {
				errorJson(w, http.StatusInternalServerError, err)
				log.Println(err)
				return
			}
			defer conn.Close()

			c := users.NewUserServiceClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			res, err := c.VerifyJWT(ctx, &users.JWTRequest{
				Token: r.Header["Token"][0],
			})
			if err != nil {
				errorJson(w, http.StatusUnauthorized, errors.New("you are not logged in"))
				log.Println(err)
				return
			}
			if res.GetResult() == false {
				errorJson(w, http.StatusUnauthorized, errors.New("you are not logged in"))
				return
			}
			endpointHandler(w, r)
		} else {
			errorJson(w, http.StatusBadRequest, errors.New("missing token"))
			return
		}
	}
}
