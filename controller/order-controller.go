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
	All(context *gin.Context)
	FindPharmacist(context *gin.Context)
	FindDoctor(context *gin.Context)
	FindOrderDoctor(context *gin.Context)
	FindOrderPharmacist(context *gin.Context)
	TranscactBonus(context *gin.Context)
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

func (c *orderController) FindOrderDoctor(context *gin.Context) {
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

	userType := fmt.Sprintf("%v", claims["user_type"])
	if userType == "doctor" {

		var orders []entity.Order = c.orderService.FindDoctor(user_id)
		if len(orders) > 0 {
			res := helper.BuildResponse(true, "OK", orders)
			context.JSON(http.StatusOK, res)
		} else {
			res := helper.BuildResponse(true, "OK", helper.EmptyObj{})
			context.JSON(http.StatusNotFound, res)
		}
	} else {
		res := helper.BuildErrorResponse("Permission denied", "Permission denied", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
}

func (c *orderController) FindOrderPharmacist(context *gin.Context) {
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

	userType := fmt.Sprintf("%v", claims["user_type"])
	if userType == "pharmacist" {

		var orders []entity.Order = c.orderService.FindPharmacist(user_id)
		if len(orders) > 0 {
			res := helper.BuildResponse(true, "OK", orders)
			context.JSON(http.StatusOK, res)
		} else {
			res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
			context.JSON(http.StatusNotFound, res)
		}
	} else {
		res := helper.BuildErrorResponse("Permission denied", "Permission denied", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
}

func (c *orderController) FindDoctor(context *gin.Context) {
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

	userType := fmt.Sprintf("%v", claims["user_type"])
	if userType == "admin" {
		id, err := strconv.ParseUint(context.Query("id"), 0, 0)
		if err != nil {
			res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		var orders []entity.Order = c.orderService.FindDoctor(id)
		if len(orders) > 0 {
			res := helper.BuildResponse(true, "OK", orders)
			context.JSON(http.StatusOK, res)
		} else {
			res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
			context.JSON(http.StatusNotFound, res)
		}
	} else {
		res := helper.BuildErrorResponse("Permission denied", "Permission denied", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}

}

func (c *orderController) FindPharmacist(context *gin.Context) {
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

	userType := fmt.Sprintf("%v", claims["user_type"])
	if userType == "admin" {
		id, err := strconv.ParseUint(context.Query("id"), 0, 0)
		if err != nil {
			res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		var orders []entity.Order = c.orderService.FindPharmacist(id)
		if len(orders) > 0 {
			res := helper.BuildResponse(true, "OK", orders)
			context.JSON(http.StatusOK, res)
		} else {
			res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
			context.JSON(http.StatusNotFound, res)
		}
	}

}

func (c *orderController) All(context *gin.Context) {
	var orders []entity.Order = c.orderService.All()
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
	if userType == "admin" {
		res := helper.BuildResponseWithCount(true, "OK", orders, len(orders))
		fmt.Println("Order count", len(orders))
		context.JSON(http.StatusOK, res)
	} else {
		res := helper.BuildErrorResponse("Permission denied", "Permission denied", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
}

func (c *orderController) TranscactBonus(context *gin.Context) {
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
	if userType == "admin" {
		var orders []entity.Order = c.orderService.TranscactBonus()
		res := helper.BuildResponseWithCount(true, "OK", orders, len(orders))
		fmt.Println("Order count", len(orders))
		context.JSON(http.StatusOK, res)
	} else {
		res := helper.BuildErrorResponse("Permission denied", "Permission denied", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
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
	fmt.Println("orderDTO : ", orderCreateDTO)
	userType := fmt.Sprintf("%v", claims["user_type"])
	// pharmacist id validate
	if userType == "pharmacist" {
		errDTO := context.ShouldBindJSON(&orderCreateDTO)
		if errDTO != nil {
			res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
			context.JSON(http.StatusBadRequest, res)

		} else {

			result, err := c.orderService.Insert(orderCreateDTO, id)
			if err == nil {
				response := helper.BuildResponse(true, "OK", result)
				context.JSON(http.StatusCreated, response)
			} else {
				response := helper.BuildErrorResponse("Permission denied", "Pharmacist id did not match", err.Error())
				context.JSON(http.StatusBadRequest, response)
			}

		}
	} else {
		response := helper.BuildErrorResponse("Permission denied", "Permission denied", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}

}
