package repository

import (
	"log"
	"time"

	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/external_api"
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
	pharm.UpdatedAt = time.Now()
	external_api.CreateLocal(pharm.Name)
	url, err := external_api.Uploader()
	if err != nil {
		log.Println(err)
		panic("Failed to upload aws")
	}

	pharm.Name = url
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

func (db *productConnection) AllPharm() []entity.Pharm {
	var pharms []entity.Pharm
	db.connection.Preload("User").Find(&pharms)
	return pharms
}
