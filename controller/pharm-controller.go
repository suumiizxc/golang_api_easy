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
	Update(context *gin.Context)
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
		// authHeader := context.GetHeader("Authorization")
		// userID := c.getUserIDByToken(authHeader)
		// convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		// if err == nil {
		// 	bookCreateDTO.UserID = convertedUserID
		// }
		result := c.pharmService.Insert(pharmCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *pharmController) Update(context *gin.Context) {
	var pharmUpdateDTO dto.PharmUpdateDTO
	errDTO := context.ShouldBind(&pharmUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	// authHeader := context.GetHeader("Authorization")
	// token, errToken := c.jwtService.ValidateToken(authHeader)
	// if errToken != nil {
	// 	panic(errToken.Error())
	// }
	// claims := token.Claims.(jwt.MapClaims)
	// userID := fmt.Sprintf("%v", claims["user_id"])
	// if c.bookService.IsAllowedToEdit(userID, bookUpdateDTO.ID) {
	// 	id, errID := strconv.ParseUint(userID, 10, 64)
	// 	if errID == nil {
	// 		bookUpdateDTO.UserID = id
	// 	}
	result := c.pharmService.Update(pharmUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
	// } else {
	// response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
	// context.JSON(http.StatusForbidden, response)
	// }
}
