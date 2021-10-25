package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/repository"
)

type ChallangeService interface {
	Insert(b entity.Challenge) entity.Challenge
}

type challengeService struct {
	challengeRepository repository.ChallengeRepository
}

func NewChallengeService(challengeRepo repository.ChallengeRepository) ChallangeService {
	return &challengeService{
		challengeRepository: challengeRepo,
	}
}

func (service *challengeService) Insert(b entity.Challenge) entity.Challenge {
	challenge := entity.Challenge{}
	err := smapping.FillStruct(&challenge, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
		return challenge
	} else {
		res := service.challengeRepository.InsertChallenge(challenge)
		return res
	}
}
