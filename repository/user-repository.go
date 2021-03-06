package repository

import (
	"log"
	"time"

	"github.com/suumiizxc/golang_api/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userID string) entity.User
	AllDoctors() []entity.Doctor
	AllPharmacist() []entity.Pharmacist
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	user.UserType = "admin"
	user.UpdatedAt = time.Now()
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	var userFind entity.User
	db.connection.Preload("User").Find(&userFind, user.ID)
	user.UserType = userFind.UserType
	user.UpdatedAt = time.Now()
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) entity.User {
	var user entity.User
	db.connection.Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

func (db *userConnection) AllDoctors() []entity.Doctor {
	var doctors []entity.Doctor
	db.connection.Preload("User").Find(&doctors)
	return doctors
}

func (db *userConnection) AllPharmacist() []entity.Pharmacist {
	var pharmacists []entity.Pharmacist
	db.connection.Preload("User").Find(&pharmacists)
	return pharmacists
}
