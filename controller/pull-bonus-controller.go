package controller

import (
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/suumiizxc/golang_api/service"
)

type PullBonusController interface {
	Insert(context *gin.Context)
	// All(context *gin.Context)
	// FindPharmacist(context *gin.Context)
	// FindDoctor(context *gin.Context)
	// FindOrderDoctor(context *gin.Context)
	// FindOrderPharmacist(context *gin.Context)
	// TranscactBonus(context *gin.Context)
}

type pullBonusController struct {
	pullBonusService service.PullBonusService
	jwtService       service.JWTService
}

//NewBookController create a new instances of BoookController
func NewPullBonusController(pullBonusServ service.PullBonusService, jwtServ service.JWTService) PullBonusController {
	return &pullBonusController{
		pullBonusService: pullBonusServ,
		jwtService:       jwtServ,
	}
}

func (c *pullBonusController) Insert(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	fmt.Println("AuthHEADER : ", authHeader)
	token, errToken := c.jwtService.ValidateToken(authHeader)

	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	user_id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	fmt.Println(user_id, "id")
	if err != nil {
		panic(err.Error())
	}

	// userType := fmt.Sprintf("%v", claims["user_type"])
	// if userType == "doctor" {

	// }
}
