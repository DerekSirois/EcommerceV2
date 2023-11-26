package Handlers

import (
	"context"
	products "gateway/pb/product/proto"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Quantity    int
	Price       float32
}

func GetAllProduct(w http.ResponseWriter, _ *http.Request) {
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

func GetAllAvailableProduct(w http.ResponseWriter, _ *http.Request) {
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

	res, err := c.GetAllAvailable(ctx, &products.Empty{})
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

func GetAllOutOfStockProduct(w http.ResponseWriter, _ *http.Request) {
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

	res, err := c.GetAllOutOfStock(ctx, &products.Empty{})
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

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	p := Product{}
	err := readJson(r, &p)
	if err != nil {
		errorJson(w, http.StatusBadRequest, err)
		log.Println(err)
		return
	}
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

	res, err := c.Create(ctx, &products.ProductRequest{ProductEntry: &products.Product{
		Name:        p.Name,
		Description: p.Description,
		Quantity:    int32(p.Quantity),
		Price:       p.Price,
	}})
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

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		errorJson(w, http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	p := Product{}
	err = readJson(r, &p)
	if err != nil {
		errorJson(w, http.StatusBadRequest, err)
		log.Println(err)
		return
	}
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

	res, err := c.Update(ctx, &products.ProductRequest{ProductEntry: &products.Product{
		Id:          int32(id),
		Name:        p.Name,
		Description: p.Description,
		Quantity:    int32(p.Quantity),
		Price:       p.Price,
	}})
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

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		errorJson(w, http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
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

	res, err := c.Delete(ctx, &products.ProductIdRequest{Id: int32(id)})
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
