package controller

import (
	"hacktiv8_fp_2/common"
	"hacktiv8_fp_2/dto"
	"hacktiv8_fp_2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController interface {
	CreateComment(ctx *gin.Context)
	GetComment(ctx *gin.Context)
	UpdateCommentByID(ctx *gin.Context)
	DeleteCommentByID(ctx *gin.Context)
}

type commentController struct {
	commentService service.CommentService
	jwtService     service.JWTService
}

func NewCommentController(cs service.CommentService, js service.JWTService) CommentController {
	return &commentController{
		commentService: cs,
		jwtService:     js,
	}
}

// CreateComment implements CommentController
func (c *commentController) CreateComment(ctx *gin.Context) {
	var commentDTO dto.CommentCreateDTO
	if err := ctx.ShouldBind(&commentDTO); err != nil {
		res := common.BuildErrorResponse("Failed to bind comment request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token := ctx.MustGet("token").(string)
	userID, _ := c.jwtService.GetUserIDByToken(token)

	commentDTO.UserID = uint64(userID)

	result, err := c.commentService.CreateComment(ctx.Request.Context(), commentDTO)
	if err != nil {
		res := common.BuildErrorResponse("Failed to add comment", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

// GetComment implements CommentController
func (c *commentController) GetComment(ctx *gin.Context) {
	result, err := c.commentService.GetComment(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to get comment", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

// UpdateCommentByID implements CommentController
/*func (c *commentController) UpdateCommentByID(ctx *gin.Context) {
	id := ctx.Param("commentID")
	commentID, _ := strconv.ParseUint(id, 10, 64)

	result, err := c.commentService.UpdateCommentByID(ctx.Request.Context(), commentID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update comment", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}*/

// DeleteCommentByID implements CommentController
func (c *commentController) DeleteCommentByID(ctx *gin.Context) {
	commentID := ctx.MustGet("commentID").(uint64)
	err := c.commentService.DeleteCommentByID(ctx.Request.Context(), commentID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to delete comment", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "OK", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
