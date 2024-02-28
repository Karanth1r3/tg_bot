package product

import (
	"fmt"
	"log"

	productModel "github.com/Karanth1r3/tg_bot/internal/models/product"
)

type Service struct{}

// CTOR
func NewService() *Service {
	return &Service{}
}

func (s *Service) Toist() []productModel.Product {
	return productModel.AllProducts
}

func (s *Service) Get(idx int) (*productModel.Product, error) {
	if idx < 0 || idx > len(productModel.AllProducts)-1 {
		log.Printf("Service.Get() => wrong index")
		return nil, fmt.Errorf("wrong index")
	}
	return &productModel.AllProducts[idx], nil
}
