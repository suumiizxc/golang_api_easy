package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/helper"
	"github.com/suumiizxc/golang_api/service"
)

type OrderController interface {
	Insert(context *gin.Context)
}

type orderController struct {
	orderService service.OrderService
	jwtService   service.JWTService
}

//NewBookController create a new instances of BoookController
func NewOrderController(orderServ service.OrderService, jwtServ service.JWTService) OrderController {
	return &orderController{
		orderService: orderServ,
		jwtService:   jwtServ,
	}
}

func (c *orderController) Insert(context *gin.Context) {
	var orderCreateDTO entity.Order
	authHeader := context.GetHeader("Authorization")
	fmt.Println("AuthHEADER : ", authHeader)
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
	userType := fmt.Sprintf("%v", claims["user_type"])

	if userType == "doctor" {
		errDTO := context.ShouldBindJSON(&orderCreateDTO)
		if errDTO != nil {
			res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
			context.JSON(http.StatusBadRequest, res)

		} else {

			result := c.orderService.Insert(orderCreateDTO)

			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusCreated, response)

		}
	} else {
		response := helper.BuildErrorResponse("Permission denied", "Permission denied", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}

}
