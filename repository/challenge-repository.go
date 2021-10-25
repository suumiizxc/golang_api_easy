package repository

import (
	"github.com/suumiizxc/golang_api/entity"
	"gorm.io/gorm"
)

type ChallengeRepository interface {
	InsertChallenge(b entity.Challenge) entity.Challenge
}

type challengeConnection struct {
	connection *gorm.DB
}

func NewChallengeRepository(dbConn *gorm.DB) ChallengeRepository {
	return &challengeConnection{
		connection: dbConn,
	}
}

func (db *challengeConnection) InsertChallenge(b entity.Challenge) entity.Challenge {
	db.connection.Save(&b)
	return b
}
