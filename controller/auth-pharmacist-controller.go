package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suumiizxc/golang_api/dto"
	"github.com/suumiizxc/golang_api/entity"
	"github.com/suumiizxc/golang_api/helper"
	"github.com/suumiizxc/golang_api/service"
)

type AuthPharmacistController interface {
	LoginPharmacist(ctx *gin.Context)
	RegisterPharmacist(ctx *gin.Context)
}

type authPharmacistController struct {
	authPharmacistService service.AuthPharmacistService
	jwtService            service.JWTService
}

func NewAuthPharmacistController(authPharmacistService service.AuthPharmacistService, jwtService service.JWTService) AuthPharmacistController {
	return &authPharmacistController{
		authPharmacistService: authPharmacistService,
		jwtService:            jwtService,
	}
}

func (c *authPharmacistController) LoginPharmacist(ctx *gin.Context) {
	var loginDTO dto.LoginPharmacistDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult := c.authPharmacistService.VerifyCredentialPharmacist(loginDTO.Email, loginDTO.Password)

	if v, ok := authResult.(entity.Pharmacist); ok {
		generateToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generateToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authPharmacistController) RegisterPharmacist(ctx *gin.Context) {
	var registerDTO dto.RegisterPharmacistDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if registerDTO.UserType == "admin" || registerDTO.UserType == "pharmacist" || registerDTO.UserType == "doctor" {
		if !c.authPharmacistService.IsDuplicateEmailPharmacist(registerDTO.Email) {
			response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
			ctx.JSON(http.StatusConflict, response)
		} else {

			createUser := c.authPharmacistService.CreatePharmacist(registerDTO)
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
