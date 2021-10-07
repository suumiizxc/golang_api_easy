package repository

import (
	"github.com/suumiizxc/golang_api/entity"
	"gorm.io/gorm"
)

type PharmRepository interface {
	InsertPharm(pharm entity.Pharm) entity.Pharm
	UpdatePharm(pharm entity.Pharm) entity.Pharm
	FindByName(name string) entity.Pharm
}

type pharmConnection struct {
	connection *gorm.DB
}

func NewPharmRepository(db *gorm.DB) PharmRepository {
	return &pharmConnection{
		connection: db,
	}
}

func (db *pharmConnection) InsertPharm(pharm entity.Pharm) entity.Pharm {
	db.connection.Save(&pharm)
	return pharm
}

func (db *pharmConnection) UpdatePharm(pharm entity.Pharm) entity.Pharm {
	db.connection.Save(&pharm)
	return pharm
}

func (db *pharmConnection) FindByName(name string) entity.Pharm {
	var pharm entity.Pharm
	db.connection.Where("name = ?", name).Take(&pharm)
	return pharm
}

func (db *bookConnection) AllPharm() []entity.Pharm {
	var pharms []entity.Pharm
	db.connection.Preload("User").Find(&pharms)
	return pharms
}
