package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/suumiizxc/golang_api/dto"
	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/repository"
)

type AuthDoctorService interface {
	VerifyCredentialDoctor(email string, password string) interface{}
	CreateDoctor(user dto.RegisterDoctorDTO) entity.Doctor
	FindByEmailDoctor(email string) entity.Doctor
	IsDuplicateEmailDoctor(email string) bool
}

type authDoctorService struct {
	doctorRepository repository.DoctorRepository
}

func NewAuthDoctorService(doctorRepo repository.DoctorRepository) AuthDoctorService {
	return &authDoctorService{
		doctorRepository: doctorRepo,
	}
}

func (service *authDoctorService) VerifyCredentialDoctor(email string, password string) interface{} {
	res := service.doctorRepository.VerifyCredentialDoctor(email, password)
	if v, ok := res.(entity.Doctor); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return true
}

func (service *authDoctorService) CreateDoctor(user dto.RegisterDoctorDTO) entity.Doctor {
	userToCreate := entity.Doctor{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.doctorRepository.InsertDoctor(userToCreate)
	return res
}

func (service *authDoctorService) FindByEmailDoctor(email string) entity.Doctor {
	return service.doctorRepository.FindByEmailDoctor(email)
}

func (service *authDoctorService) IsDuplicateEmailDoctor(email string) bool {
	res := service.doctorRepository.IsDuplicateEmailDoctor(email)
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
