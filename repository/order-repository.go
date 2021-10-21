package repository

import (
	"encoding/json"
	"fmt"

	"github.com/suumiizxc/golang_api/entity"
	"gorm.io/gorm"
)

type OrderRepository interface {
	InsertOrder(b entity.Order) entity.Order
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

func (db *orderConnection) InsertOrder(b entity.Order) entity.Order {
	// fmt.Println("Repo:", b.List)
	var bird []Bird
	b.Status = "pending"
	json.Unmarshal([]byte(b.List), &bird)
	var total_price float64

	for i, v := range bird {
		fmt.Println("ID :", i, " object:", v.Product_ID)
		var product entity.Product
		db.connection.Preload("User").Find(&product, v.Product_ID)
		total_price = total_price + float64(product.Price)*float64(v.Quantity)
	}
	b.TotalPrice = total_price

	db.connection.Save(&b)
	return b
}
