package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/suumiizxc/golang_api/dto"
	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/repository"
)

type ProductService interface {
	Insert(b dto.ProductCreateDTO) entity.Product
	Update(b dto.ProductUpdateDTO) entity.Product
	Delete(b entity.Product)
	All() []entity.Product
	FindByID(bookID uint64) entity.Product
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepo,
	}
}

func (service *productService) Insert(b dto.ProductCreateDTO) entity.Product {
	product := entity.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.productRepository.InsertProduct(product)
	return res
}

func (service *productService) Update(b dto.ProductUpdateDTO) entity.Product {
	product := entity.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.productRepository.UpdateProduct(product)
	return res
}

func (service *productService) Delete(b entity.Product) {
	service.productRepository.DeleteProduct(b)
}

func (service *productService) All() []entity.Product {
	return service.productRepository.AllProduct()
}

func (service *productService) FindByID(productID uint64) entity.Product {
	return service.productRepository.FindByProductID(productID)
}

func (service *productService) IsAllowedToEdit(userID string, bookID uint64) bool {
	b := service.productRepository.FindByProductID(bookID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
