package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/suumiizxc/golang_api/dto"
	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/helper"
	"github.com/suumiizxc/golang_api/service"
)

//BookController is a ...
type ProductController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type productController struct {
	productService service.ProductService
	jwtService     service.JWTService
}

//NewBookController create a new instances of BoookController
func NewBookController(productServ service.ProductService, jwtServ service.JWTService) ProductController {
	return &productController{
		productService: productServ,
		jwtService:     jwtServ,
	}
}

func (c *productController) All(context *gin.Context) {

	var products []entity.Product = c.productService.All()
	res := helper.BuildResponseWithCount(true, "OK", products, len(products))
	fmt.Println("Books count", len(products))
	context.JSON(http.StatusOK, res)
}

func (c *productController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var product entity.Product = c.productService.FindByID(id)
	if (product == entity.Product{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", product)
		context.JSON(http.StatusOK, res)
	}
}

func (c *productController) Insert(context *gin.Context) {
	var productCreateDTO dto.ProductCreateDTO
	errDTO := context.ShouldBind(&productCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID, userType := c.getUserIdUserTypeByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			productCreateDTO.UserID = convertedUserID
		}
		if userType == "admin" {
			result := c.productService.Insert(productCreateDTO)
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusCreated, response)
		} else {
			response := helper.BuildErrorResponse("Permission denied", "Permission denied", helper.EmptyObj{})
			context.JSON(http.StatusBadRequest, response)
		}

	}
}

func (c *productController) Update(context *gin.Context) {
	var productUpdateDTO dto.ProductUpdateDTO
	errDTO := context.ShouldBind(&productUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	userType := fmt.Sprintf("%v", claims["user_type"])
	if userType == "admin" {
		if c.productService.IsAllowedToEdit(userID, productUpdateDTO.ID) {
			id, errID := strconv.ParseUint(userID, 10, 64)
			if errID == nil {
				productUpdateDTO.UserID = id
			}
			result := c.productService.Update(productUpdateDTO)
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusOK, response)
		} else {
			response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
			context.JSON(http.StatusForbidden, response)
		}
	} else {
		response := helper.BuildErrorResponse("Permission denied", "Permission denied", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}

}

func (c *productController) Delete(context *gin.Context) {
	var product entity.Product
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	product.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	userType := fmt.Sprintf("%v", claims["user_type"])
	if userType == "admin" {
		if c.productService.IsAllowedToEdit(userID, product.ID) {
			c.productService.Delete(product)
			res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
			context.JSON(http.StatusOK, res)
		} else {
			response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
			context.JSON(http.StatusForbidden, response)
		}
	} else {
		response := helper.BuildErrorResponse("Permission denied", "Permission denied", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}

}

func (c *productController) getUserIdUserTypeByToken(token string) (string, string) {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user_type := fmt.Sprintf("%v", claims["user_type"])
	return id, user_type
}
