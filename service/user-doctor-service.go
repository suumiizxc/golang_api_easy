package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/suumiizxc/golang_api/dto"
	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/repository"
)

type DoctorService interface {
	UpdateDoctor(user dto.DoctorUpdateDTO) entity.Doctor
	ProfileDoctor(userID string) entity.Doctor
	AllDoctorsOrderList() []entity.APIOrderList
}

type doctorService struct {
	doctorRepository repository.DoctorRepository
}

func NewDoctorService(doctorRepo repository.DoctorRepository) DoctorService {
	return &doctorService{
		doctorRepository: doctorRepo,
	}
}

func (service *doctorService) AllDoctorsOrderList() []entity.APIOrderList {
	return service.doctorRepository.AllDoctorsOrderList()
}

func (service *doctorService) UpdateDoctor(user dto.DoctorUpdateDTO) entity.Doctor {
	userToUpdate := entity.Doctor{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.doctorRepository.UpdateDoctor(userToUpdate)
	return updatedUser
}

func (service *doctorService) ProfileDoctor(userID string) entity.Doctor {
	return service.doctorRepository.ProfileDoctor(userID)
}
