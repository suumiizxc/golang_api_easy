package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suumiizxc/golang_api/dto"
	"github.com/suumiizxc/golang_api/helper"
	"github.com/suumiizxc/golang_api/service"
)

//BookController is a ...
type PharmController interface {
	Insert(context *gin.Context)
}

type pharmController struct {
	pharmService service.PharmService
}

//NewBookController create a new instances of BoookController
func NewPharmController(pharmServ service.PharmService) PharmController {
	return &pharmController{
		pharmService: pharmServ,
	}
}

func (c *pharmController) Insert(context *gin.Context) {
	var pharmCreateDTO dto.PharmCreateDTO
	errDTO := context.ShouldBind(&pharmCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {

		result := c.pharmService.Insert(pharmCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}
