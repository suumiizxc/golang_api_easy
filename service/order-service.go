package service

import (
	"errors"
	"log"

	"github.com/mashingan/smapping"
	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/repository"
)

type OrderService interface {
	Insert(b entity.Order, tokenID uint64) (entity.Order, error)
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepo,
	}
}

func (service *orderService) Insert(b entity.Order, tokenID uint64) (entity.Order, error) {
	order := entity.Order{}
	// product := entity.Product{}
	err := smapping.FillStruct(&order, smapping.MapFields(&b))
	if order.PharmacistID == tokenID {
		if err != nil {
			log.Fatalf("Failed map %v", err)
		}
		res := service.orderRepository.InsertOrder(order)
		return res, nil
	} else {
		return order, errors.New("empty name")
	}

}
