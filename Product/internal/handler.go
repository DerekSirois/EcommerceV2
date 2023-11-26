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

func (ps *ProductServer) GetAllAvailable(_ context.Context, _ *products.Empty) (*products.ProductListResponse, error) {
	p, err := db.GetAllAvailable()
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

func (ps *ProductServer) GetAllOutOfStock(_ context.Context, _ *products.Empty) (*products.ProductListResponse, error) {
	p, err := db.GetAllOutOfStock()
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

func (ps *ProductServer) GetById(_ context.Context, pId *products.ProductIdRequest) (*products.ProductSingleResponse, error) {
	id := int(pId.GetId())
	p, err := db.GetById(id)
	if err != nil {
		return &products.ProductSingleResponse{}, err
	}

	pSingle := products.ProductSingleResponse{
		Products: &products.Product{
			Id:          int32(p.Id),
			Name:        p.Name,
			Description: p.Description,
			Quantity:    int32(p.Quantity),
			Price:       p.Price,
		},
	}

	return &pSingle, nil
}

func (ps *ProductServer) Create(_ context.Context, r *products.ProductRequest) (*products.ProductResponse, error) {
	input := r.GetProductEntry()

	err := db.Create(db.Product{
		Name:        input.Name,
		Description: input.Description,
		Quantity:    int(input.Quantity),
		Price:       input.Price,
	})
	if err != nil {
		return &products.ProductResponse{Result: "Error creating in the database"}, err
	}

	return &products.ProductResponse{Result: "Product created successfully"}, err
}

func (ps *ProductServer) Update(_ context.Context, r *products.ProductRequest) (*products.ProductResponse, error) {
	input := r.GetProductEntry()

	err := db.Update(db.Product{
		Name:        input.Name,
		Description: input.Description,
		Quantity:    int(input.Quantity),
		Price:       input.Price,
	}, int(input.Id))
	if err != nil {
		return &products.ProductResponse{Result: "Error updating in the database"}, err
	}

	return &products.ProductResponse{Result: "Product updating successfully"}, err
}

func (ps *ProductServer) Delete(_ context.Context, r *products.ProductIdRequest) (*products.ProductResponse, error) {
	id := r.GetId()

	err := db.Delete(int(id))
	if err != nil {
		return &products.ProductResponse{Result: "Error deleting in the database"}, err
	}

	return &products.ProductResponse{Result: "Product deleted successfully"}, err
}
