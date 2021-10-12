package repository

import (
	"log"
	"time"

	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/external_api"
	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProduct(b entity.Product) entity.Product
	UpdateProduct(b entity.Product) entity.Product
	DeleteProduct(b entity.Product)
	AllProduct() []entity.Product
	FindByProductID(productID uint64) entity.Product
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(dbConn *gorm.DB) ProductRepository {
	return &productConnection{
		connection: dbConn,
	}
}

func (db *productConnection) InsertProduct(b entity.Product) entity.Product {
	db.connection.Save(&b)
	b.UpdatedAt = time.Now()
	external_api.CreateLocal(b.Image)
	url, err := external_api.Uploader()
	if err != nil {
		log.Println(err)
		panic("Failed to upload aws")
	}

	b.Image = url
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *productConnection) UpdateProduct(b entity.Product) entity.Product {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *productConnection) DeleteProduct(b entity.Product) {
	db.connection.Delete(&b)
}

func (db *productConnection) FindByProductID(productID uint64) entity.Product {
	var product entity.Product
	db.connection.Preload("User").Find(&product, productID)
	return product
}

func (db *productConnection) AllProduct() []entity.Product {
	var products []entity.Product
	db.connection.Preload("User").Find(&products)
	return products
}
