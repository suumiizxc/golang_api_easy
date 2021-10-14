package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suumiizxc/golang_api/dto"
	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/helper"
	"github.com/suumiizxc/golang_api/service"
)

type AuthDoctorController interface {
	LoginDoctor(ctx *gin.Context)
	RegisterDoctor(ctx *gin.Context)
}

type authDoctorController struct {
	authDoctorService service.AuthDoctorService
	jwtService        service.JWTService
}

func NewAuthDoctorController(authDoctorService service.AuthDoctorService, jwtService service.JWTService) AuthDoctorController {
	return &authDoctorController{
		authDoctorService: authDoctorService,
		jwtService:        jwtService,
	}
}

func (c *authDoctorController) LoginDoctor(ctx *gin.Context) {
	var loginDTO dto.LoginDoctorDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authDoctorService.VerifyCredentialDoctor(loginDTO.Email, loginDTO.Password)
	fmt.Println("AuthResultDoctor: ", authResult)
	if v, ok := authResult.(entity.Doctor); ok {
		generateToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generateToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authDoctorController) RegisterDoctor(ctx *gin.Context) {
	var registerDTO dto.RegisterDoctorDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if registerDTO.UserType == "admin" || registerDTO.UserType == "pharmacist" || registerDTO.UserType == "doctor" {
		if !c.authDoctorService.IsDuplicateEmailDoctor(registerDTO.Email) {
			response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
			ctx.JSON(http.StatusConflict, response)
		} else {

			createUser := c.authDoctorService.CreateDoctor(registerDTO)
			token := c.jwtService.GenerateToken(strconv.FormatUint(createUser.ID, 10))
			createUser.Token = token
			response := helper.BuildResponse(true, "OK!", createUser)
			ctx.JSON(http.StatusCreated, response)

		}
	} else {
		response := helper.BuildErrorResponse("Failed to process", "User type not match", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}

}
