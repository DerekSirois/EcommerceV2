package Handlers

import (
	"context"
	products "gateway/pb/product/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
)

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("product:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		errorJson(w, http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	defer conn.Close()

	c := products.NewProductServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.GetAll(ctx, &products.Empty{})
	if err != nil {
		errorJson(w, http.StatusInternalServerError, err)
		log.Println(err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Get all",
		Data:    res.GetProducts(),
	}

	_ = sendJson(w, http.StatusOK, payload)
}
