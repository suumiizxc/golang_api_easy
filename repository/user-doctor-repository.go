package repository

import (
	"time"

	"github.com/suumiizxc/golang_api/entity"
	"gorm.io/gorm"
)

type DoctorRepository interface {
	InsertDoctor(user entity.Doctor) entity.Doctor
	UpdateDoctor(user entity.Doctor) entity.Doctor
	VerifyCredentialDoctor(email string, password string) interface{}
	IsDuplicateEmailDoctor(email string) (tx *gorm.DB)
	FindByEmailDoctor(email string) entity.Doctor
	ProfileDoctor(userID string) entity.Doctor
}

type doctorConnection struct {
	connection *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) DoctorRepository {
	return &doctorConnection{
		connection: db,
	}
}

func (db *doctorConnection) InsertDoctor(user entity.Doctor) entity.Doctor {
	user.Password = hashAndSalt([]byte(user.Password))

	user.UpdatedAt = time.Now()
	db.connection.Save(&user)
	return user
}

func (db *doctorConnection) UpdateDoctor(user entity.Doctor) entity.Doctor {
	user.Password = hashAndSalt([]byte(user.Password))
	user.UpdatedAt = time.Now()
	db.connection.Save(&user)
	return user
}

func (db *doctorConnection) VerifyCredentialDoctor(email string, password string) interface{} {
	var user entity.Doctor
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *doctorConnection) IsDuplicateEmailDoctor(email string) (tx *gorm.DB) {
	var user entity.Doctor
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *doctorConnection) FindByEmailDoctor(email string) entity.Doctor {
	var user entity.Doctor
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *doctorConnection) ProfileDoctor(userID string) entity.Doctor {
	var user entity.Doctor
	db.connection.Find(&user, userID)
	return user
}
