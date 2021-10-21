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

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	CheckToken(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	fmt.Println("AuthResult :", authResult)
	if v, ok := authResult.(entity.User); ok {
		generateToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10), "admin")
		v.Token = generateToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {

		createUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createUser.ID, 10), "admin")
		createUser.Token = token
		response := helper.BuildResponse(true, "OK!", createUser)
		ctx.JSON(http.StatusCreated, response)

	}

}

func (c *authController) CheckToken(context *gin.Context) {
	var checkTokenDTO dto.CheckTokenDTO
	errDTO := context.ShouldBind(&checkTokenDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	token, err := c.jwtService.ValidateToken(checkTokenDTO.Token)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	user_type := fmt.Sprintf("%v", claims["user_type"])
	if user_type == "admin" {
		res := helper.BuildResponse(true, "admin", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else if user_type == "doctor" {
		res := helper.BuildResponse(true, "doctor", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else if user_type == "pharmacist" {
		res := helper.BuildResponse(true, "pharmacist", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		res := helper.BuildErrorResponse("Not found user type in token", "Not found user type in token", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
}
