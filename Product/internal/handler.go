package internal

import (
	"context"
	"product/internal/db"
	products "product/pb/proto"
)

type ProductServer struct {
	products.UnimplementedProductServiceServer
}

func (ps *ProductServer) GetAll(_ context.Context, _ *products.Empty) (*products.ProductListResponse, error) {
	p, err := db.GetAll()
	if err != nil {
		return &products.ProductListResponse{}, err
	}

	pList := products.ProductListResponse{
		Products: make([]*products.Product, 0),
	}

	for _, item := range p {
		pList.Products = append(pList.Products, &products.Product{
			Id:          int32(item.Id),
			Name:        item.Name,
			Description: item.Description,
			Quantity:    int32(item.Quantity),
			Price:       item.Price,
		})
	}

	return &pList, nil
}
