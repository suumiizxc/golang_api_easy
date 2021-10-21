package controller

import (
	"net/http"

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
}

//NewBookController create a new instances of BoookController
func NewOrderController(orderServ service.OrderService) OrderController {
	return &orderController{
		orderService: orderServ,
	}
}

func (c *orderController) Insert(context *gin.Context) {
	var orderCreateDTO entity.Order

	errDTO := context.ShouldBindJSON(&orderCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)

	} else {

		result := c.orderService.Insert(orderCreateDTO)

		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)

	}
}
