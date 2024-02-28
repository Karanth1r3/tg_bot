package product

import productModel "github.com/Karanth1r3/tg_bot/internal/models/product"

type Service struct{}

// CTOR
func NewService() *Service {
	return &Service{}
}

func (s *Service) Toist() []productModel.Product {
	return productModel.AllProducts
}
