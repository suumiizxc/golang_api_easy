package repository

import (
	"github.com/suumiizxc/golang_api/entity"
	"gorm.io/gorm"
)

type PullBonusRepository interface {
	InsertPullBonus(b entity.PullBonus) entity.PullBonus
	// AllOrder() []entity.Order
	// DoctorOrders(page uint64, paginition uint64, ordering string) []entity.Order
	// FindByPharmacistOrder(pharmacistID uint64) []entity.Order
	// FindByDoctorOrder(doctorID uint64) []entity.Order
	// FindByOrderID(orderID uint64) entity.Order
	// TranscactBonus() []entity.Order
}

type pullBonusConnection struct {
	connection *gorm.DB
}

func NewPullBonusRepository(dbConn *gorm.DB) PullBonusRepository {
	return &pullBonusConnection{
		connection: dbConn,
	}
}

func (db *pullBonusConnection) InsertPullBonus(b entity.PullBonus) entity.PullBonus {
	// var pullBonus entity.PullBonus
	db.connection.Save(&b)
	return b
}
