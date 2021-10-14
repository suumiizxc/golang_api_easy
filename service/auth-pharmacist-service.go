package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/suumiizxc/golang_api/dto"
	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/repository"
)

type AuthPharmacistService interface {
	VerifyCredentialPharmacist(email string, password string) interface{}
	CreatePharmacist(user dto.RegisterPharmacistDTO) entity.Pharmacist
	FindByEmailPharmacist(email string) entity.Pharmacist
	IsDuplicateEmailPharmacist(email string) bool
}

type authPharmacistService struct {
	pharmacistRepository repository.PharmacistRepository
}

func NewAuthPharmacistService(pharmacistRepo repository.PharmacistRepository) AuthPharmacistService {
	return &authPharmacistService{
		pharmacistRepository: pharmacistRepo,
	}
}

func (service *authPharmacistService) VerifyCredentialPharmacist(email string, password string) interface{} {
	res := service.pharmacistRepository.VerifyCredentialPharmacist(email, password)
	if v, ok := res.(entity.Doctor); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return true
}

func (service *authPharmacistService) CreatePharmacist(user dto.RegisterPharmacistDTO) entity.Pharmacist {
	userToCreate := entity.Pharmacist{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.pharmacistRepository.InsertPharmacist(userToCreate)
	return res
}

func (service *authPharmacistService) FindByEmailPharmacist(email string) entity.Pharmacist {
	return service.pharmacistRepository.FindByEmailPharmacist(email)
}

func (service *authPharmacistService) IsDuplicateEmailPharmacist(email string) bool {
	res := service.pharmacistRepository.IsDuplicateEmailPharmacist(email)
	return !(res.Error == nil)
}

// func comparePassword(hashedPwd string, plainPassword []byte) bool {
// 	byteHash := []byte(hashedPwd)
// 	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)

// 	if err != nil {
// 		log.Println(err)
// 		return false
// 	}
// 	return true
// }
