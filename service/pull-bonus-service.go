package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/repository"
)

type PullBonusService interface {
	Insert(b entity.PullBonus) (entity.PullBonus, error)
	// All() []entity.Order
	// FindPharmacist(pharmacistID uint64) []entity.Order
	// FindDoctor(doctorID uint64) []entity.Order
	// FindOrderByID(orderID uint64) entity.Order
	// TranscactBonus() []entity.Order
}

type pullBonusService struct {
	pullBonusRepository repository.PullBonusRepository
}

func NewPullBonusService(pullBonusRepo repository.PullBonusRepository) PullBonusService {
	return &pullBonusService{
		pullBonusRepository: pullBonusRepo,
	}
}

func (service *pullBonusService) Insert(b entity.PullBonus) (entity.PullBonus, error) {
	pullBonus := entity.PullBonus{}
	// product := entity.Product{}
	err := smapping.FillStruct(&pullBonus, smapping.MapFields(&b))

	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.pullBonusRepository.InsertPullBonus(pullBonus)
	return res, nil

}
