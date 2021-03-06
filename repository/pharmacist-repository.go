package repository

import (
	"time"

	"github.com/suumiizxc/golang_api/entity"
	"gorm.io/gorm"
)

type PharmacistRepository interface {
	InsertPharmacist(user entity.Pharmacist) entity.Pharmacist
	UpdatePharmacist(user entity.Pharmacist) entity.Pharmacist
	VerifyCredentialPharmacist(email string, password string) interface{}
	IsDuplicateEmailPharmacist(email string) (tx *gorm.DB)
	FindByEmailPharmacist(email string) entity.Pharmacist
	ProfilePharmacist(userID string) entity.Pharmacist
	AllPharmacistOrderList() []entity.APIOrderList
}

type pharmacistConnection struct {
	connection *gorm.DB
}

func NewPharmacistRepository(db *gorm.DB) PharmacistRepository {
	return &pharmacistConnection{
		connection: db,
	}
}

func (db *pharmacistConnection) AllPharmacistOrderList() []entity.APIOrderList {
	var pharmacists []entity.APIOrderList
	db.connection.Model(&entity.Pharmacist{}).Order("claimed_point desc").Find(&pharmacists)
	// db.connection.Model(&entity.Doctor{}).Delete(&entity.Doctor{}, 3)
	return pharmacists
}

func (db *pharmacistConnection) InsertPharmacist(user entity.Pharmacist) entity.Pharmacist {
	user.Password = hashAndSalt([]byte(user.Password))
	user.UserType = "pharmacist"
	user.UpdatedAt = time.Now()
	user.ClaimedPoint = 0
	db.connection.Save(&user)
	return user
}

func (db *pharmacistConnection) UpdatePharmacist(user entity.Pharmacist) entity.Pharmacist {
	user.Password = hashAndSalt([]byte(user.Password))
	var userFind entity.Pharmacist
	db.connection.Preload("User").Find(&userFind, user.ID)
	user.UserType = userFind.UserType
	// user.Balance = userFind.Balance
	user.UpdatedAt = time.Now()
	db.connection.Save(&user)
	return user
}

func (db *pharmacistConnection) VerifyCredentialPharmacist(email string, password string) interface{} {
	var user entity.Pharmacist
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *pharmacistConnection) IsDuplicateEmailPharmacist(email string) (tx *gorm.DB) {
	var user entity.Pharmacist
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *pharmacistConnection) FindByEmailPharmacist(email string) entity.Pharmacist {
	var user entity.Pharmacist
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *pharmacistConnection) ProfilePharmacist(userID string) entity.Pharmacist {
	var user entity.Pharmacist
	db.connection.Find(&user, userID)
	return user
}
