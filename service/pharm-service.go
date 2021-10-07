package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/suumiizxc/golang_api/dto"
	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/repository"
)

type PharmService interface {
	Update(user dto.PharmUpdateDTO) entity.Pharm
	Insert(user dto.PharmCreateDTO) entity.Pharm
}

type pharmService struct {
	pharmRepository repository.PharmRepository
}

func NewPharmService(pharmRepo repository.PharmRepository) PharmService {
	return &pharmService{
		pharmRepository: pharmRepo,
	}
}

func (service *pharmService) Insert(b dto.PharmCreateDTO) entity.Pharm {
	pharm := entity.Pharm{}
	err := smapping.FillStruct(&pharm, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.pharmRepository.InsertPharm(pharm)
	return res
}

func (service *pharmService) Update(b dto.PharmUpdateDTO) entity.Pharm {
	pharm := entity.Pharm{}
	err := smapping.FillStruct(&pharm, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.pharmRepository.UpdatePharm(pharm)
	return res
}
