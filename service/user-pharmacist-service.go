package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/suumiizxc/golang_api/dto"
	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/repository"
)

type PharmacistService interface {
	UpdatePharmacist(user dto.PharmacistUpdateDTO) entity.Pharmacist
	ProfilePharmacist(userID string) entity.Pharmacist
	AllPharmacistOrderList() []entity.APIOrderList
}

type pharmacistService struct {
	pharmacistRepository repository.PharmacistRepository
}

func NewPharmacistService(pharmacistRepo repository.PharmacistRepository) PharmacistService {
	return &pharmacistService{
		pharmacistRepository: pharmacistRepo,
	}
}

func (service *pharmacistService) AllPharmacistOrderList() []entity.APIOrderList {
	return service.pharmacistRepository.AllPharmacistOrderList()
}

func (service *pharmacistService) UpdatePharmacist(user dto.PharmacistUpdateDTO) entity.Pharmacist {
	userToUpdate := entity.Pharmacist{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.pharmacistRepository.UpdatePharmacist(userToUpdate)
	return updatedUser
}

func (service *pharmacistService) ProfilePharmacist(userID string) entity.Pharmacist {
	return service.pharmacistRepository.ProfilePharmacist(userID)
}
