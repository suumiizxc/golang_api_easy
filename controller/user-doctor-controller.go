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

type DoctorController interface {
	UpdateDoctor(context *gin.Context)
	ProfileDoctor(context *gin.Context)
}

type doctorController struct {
	doctorService service.DoctorService
	jwtService    service.JWTService
}

func NewDoctorController(userService service.DoctorService, jwtService service.JWTService) DoctorController {
	return &doctorController{
		doctorService: userService,
		jwtService:    jwtService,
	}
}

func (c *doctorController) UpdateDoctor(context *gin.Context) {
	var userUpdateDTO dto.DoctorUpdateDTO
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
	u := c.doctorService.UpdateDoctor(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)

	context.JSON(http.StatusOK, res)
}

func (c *doctorController) ProfileDoctor(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)

	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.doctorService.ProfileDoctor(id)
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}
