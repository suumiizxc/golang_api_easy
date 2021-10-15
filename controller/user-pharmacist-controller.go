package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/suumiizxc/golang_api/dto"
	"github.com/suumiizxc/golang_api/helper"
	"github.com/suumiizxc/golang_api/service"
)

type PharmacistController interface {
	UpdatePharmacist(context *gin.Context)
	ProfilePharmacist(context *gin.Context)
}

type pharmacistController struct {
	pharmacistService service.PharmacistService
	jwtService        service.JWTService
}

func NewPharmacistController(pharmacistService service.PharmacistService, jwtService service.JWTService) PharmacistController {
	return &pharmacistController{
		pharmacistService: pharmacistService,
		jwtService:        jwtService,
	}
}

func (c *pharmacistController) UpdatePharmacist(context *gin.Context) {
	var userUpdateDTO dto.PharmacistUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)

	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)

	fmt.Println(id, "id")
	if err != nil {
		panic(err.Error())
	}

	userUpdateDTO.ID = id
	u := c.pharmacistService.UpdatePharmacist(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)

	context.JSON(http.StatusOK, res)
}

func (c *pharmacistController) ProfilePharmacist(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)

	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.pharmacistService.ProfilePharmacist(id)
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}
