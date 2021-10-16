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

type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
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
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)

	context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)

	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user_type := fmt.Sprintf("%v", claims["user_type"])
	if user_type == "admin" {
		user := c.userService.Profile(id)
		res := helper.BuildResponse(true, "OK!", user)
		context.JSON(http.StatusOK, res)
	} else {
		res := helper.BuildErrorResponse("Permission denied", "Permission denied", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}

}
