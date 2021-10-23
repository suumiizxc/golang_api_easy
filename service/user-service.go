package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/suumiizxc/golang_api/dto"
	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/repository"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
	AllDoctors() []entity.Doctor
	AllPharmacist() []entity.Pharmacist
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}

func (service *userService) AllDoctors() []entity.Doctor {
	return service.userRepository.AllDoctors()
}

func (service *userService) AllPharmacist() []entity.Pharmacist {
	return service.userRepository.AllPharmacist()
}
