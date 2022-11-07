package controller

import (
	"fmt"
	"hacktiv8_fp_2/common"
	"hacktiv8_fp_2/dto"
	"hacktiv8_fp_2/service"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocmedController interface {
	PostSocmed(ctx *gin.Context)
	GetSocmed(ctx *gin.Context)
	UpdateSocmedById(ctx *gin.Context)
	DeleteSocmedById(ctx *gin.Context)
}

type socmedController struct {
	userService   service.UserService
	jwtService    service.JWTService
	socmedService service.SocmedService
}

func NewSocmedController(us service.UserService, ss service.SocmedService, js service.JWTService) SocmedController {
	return &socmedController{
		userService:   us,
		socmedService: ss,
		jwtService:    js,
	}
}

func (sc *socmedController) PostSocmed(ctx *gin.Context) {

	body, _ := ioutil.ReadAll(ctx.Request.Body);
	fmt.Println("bodynya -> ", string(body));
	var socmedDTO dto.SocialMediaCreateDTO
	err := ctx.ShouldBind(&socmedDTO)
	if err != nil {
		response := common.BuildErrorResponse("invalid input", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Print("ini bind -> ")
	fmt.Printf("%+v\n", socmedDTO)

	token := ctx.MustGet("token").(string)
	userId, err := sc.jwtService.GetUserIDByToken(token)
	fmt.Println("User id -> ", userId)
	socmedDTO.UserID = uint64(userId)
	fmt.Println("id dto -> ", socmedDTO.UserID)

	if err != nil {
		response := common.BuildErrorResponse("invalid token", err.Error(), nil)
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	result, err := sc.socmedService.AddNewSocmed(ctx.Request.Context(), socmedDTO) // implement func ...

	if err != nil {
		response := common.BuildErrorResponse("invalid input", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := common.BuildResponse(true, "Created", result)
	ctx.JSON(http.StatusCreated, response)
}

func (sc *socmedController) GetSocmed(ctx *gin.Context)        {}
func (sc *socmedController) UpdateSocmedById(ctx *gin.Context) {}
func (sc *socmedController) DeleteSocmedById(ctx *gin.Context) {}
