package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/suumiizxc/golang_api/entity"
	"gorm.io/gorm"
)

type OrderRepository interface {
	InsertOrder(b entity.Order) entity.Order
	AllOrder() []entity.Order
	FindByPharmacistOrder(pharmacistID uint64) []entity.Order
	FindByDoctorOrder(doctorID uint64) []entity.Order
}

type orderConnection struct {
	connection *gorm.DB
}

func NewOrderRepository(dbConn *gorm.DB) OrderRepository {
	return &orderConnection{
		connection: dbConn,
	}
}

type Bird struct {
	Product_ID float64
	Quantity   float64
}

func (db *orderConnection) AllOrder() []entity.Order {
	var orders []entity.Order
	db.connection.Preload("Orders").Find(&orders)
	return orders
}

func (db *orderConnection) FindByPharmacistOrder(pharmacistID uint64) []entity.Order {
	var orders []entity.Order
	db.connection.Preload("Pharmacist").Preload("Doctor").Where("pharmacist_id = ?", pharmacistID).Find(&orders)
	return orders
}

func (db *orderConnection) FindByDoctorOrder(doctorID uint64) []entity.Order {
	var orders []entity.Order
	db.connection.Preload("Pharmacist").Preload("Doctor").Where("doctor_id = ?", doctorID).Find(&orders)
	return orders
}

func (db *orderConnection) InsertOrder(b entity.Order) entity.Order {
	var doctorC entity.Doctor
	var pharmacistC entity.Pharmacist
	var doctorS entity.Doctor
	var pharmacistS entity.Pharmacist
	doctor_coupon_rate, err := strconv.ParseFloat(os.Getenv("DOCTOR_COUPON_RATE"), 64)
	if err != nil {
		panic(err.Error())
	}
	pharmacist_coupon_rate, err := strconv.ParseFloat(os.Getenv("PHARMACIST_COUPON_RATE"), 64)
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println("Repo:", b.List)
	var bird []Bird
	b.Status = "pending"
	json.Unmarshal([]byte(b.List), &bird)
	var total_price float64
	var doctor_coupon float64
	var pharmacist_coupon float64
	for i, v := range bird {
		fmt.Println("ID :", i, " object:", v.Product_ID)
		var product entity.Product
		db.connection.Preload("User").Find(&product, v.Product_ID)
		total_price = total_price + float64(product.Price)*float64(v.Quantity)
		doctor_coupon = doctor_coupon + float64(product.Price)*doctor_coupon_rate*float64(v.Quantity)
		pharmacist_coupon = pharmacist_coupon + float64(product.Price)*pharmacist_coupon_rate*float64(v.Quantity)
	}
	b.TotalPrice = total_price
	b.CouponDoctor = doctor_coupon
	b.CouponPharmacist = pharmacist_coupon

	db.connection.Model(&doctorS).Find(&doctorS, b.DoctorID)
	db.connection.Model(&pharmacistS).Find(&pharmacistS, b.PharmacistID)
	fmt.Println("DOCTORS : ", doctorS)

	db.connection.Model(&doctorC).Where("id = ?", b.DoctorID).Update("balance", (doctor_coupon + doctorS.Balance))
	db.connection.Model(&pharmacistC).Where("id = ?", b.PharmacistID).Update("balance", (pharmacist_coupon + pharmacistS.Balance))

	db.connection.Save(&b)
	db.connection.Preload("Pharmacist").Preload("Doctor").Find(&b)
	return b
}
