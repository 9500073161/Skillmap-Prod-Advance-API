package handlers

import (
	"fmt"
	"net/http"

	"github.com/9500073161/skill-map-prod/common"
	"github.com/9500073161/skill-map-prod/managers"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	groupName   string
	userManager managers.UserManager
}

func NewUserHandlerFrom(userManager managers.UserManager) *UserHandler {
	return &UserHandler{
		"api/users/",
		userManager,
	}
}

func (handler *UserHandler) RegisterUserApis(r *gin.Engine) {
	userGroup := r.Group(handler.groupName)

	userGroup.GET("", handler.List)
	userGroup.POST("", handler.Create)
	userGroup.GET(":userid/", handler.Detail)
	userGroup.DELETE(":userid/", handler.Delete)
	userGroup.PATCH(":userid/", handler.Update)
}

func (handler *UserHandler) Create(ctx *gin.Context) {

	userData := common.NewUserCreationInput()

	err := ctx.BindJSON(&userData)

	if err != nil {
		common.BadResponse(ctx, "Failed to bind data")
		return
	}

	newUser, err := handler.userManager.Create(userData)

	if err != nil {
		common.BadResponse(ctx, "failed to create user")
		return
	}

	ctx.JSON(http.StatusOK, newUser)
}

func (handler *UserHandler) List(ctx *gin.Context) {

	allUsers, err := handler.userManager.List()

	if err != nil {
		common.BadResponse(ctx, "failed to get users")
		return
	}

	ctx.JSON(http.StatusOK, allUsers)
}

func (handler *UserHandler) Detail(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userid")

	if !ok {
		fmt.Println("invalid userid")
	}
	user, err := handler.userManager.Get(userId)

	if user.ID == 0 {
		common.BadResponse(ctx, "no user present")
		return
	}

	if err != nil {
		common.BadResponse(ctx, "failed to get user")
	}

	ctx.JSON(http.StatusOK, user)
}

func (handler *UserHandler) Delete(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userid")

	if !ok {
		common.BadResponse(ctx, "invalid userid")
	}
	err := handler.userManager.Delete(userId)

	if err != nil {
		common.BadResponse(ctx, "failed to delete user")
	}

	common.SuccessResponse(ctx, "Deleted user")
}

func (handler *UserHandler) Update(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userid")

	if !ok {
		common.BadResponse(ctx, "failed to delete user")
	}

	userData := common.NewUserUpdateInput()

	err := ctx.BindJSON(&userData)

	if err != nil {
		common.BadResponse(ctx, "failed to bind data")
		return
	}

	user, err := handler.userManager.Update(userId, userData)

	if err != nil {
		common.BadResponse(ctx, "failed to update user")
		return
	}

	ctx.JSON(http.StatusOK, user)
}
