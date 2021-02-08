package handlers

import (
	"github.com/gin-gonic/gin"
	user "problem/1/build/gen/api"
	"problem/1/pkg/models"
)

type UserDetailParams struct {
	ID int32 `uri:"id" binding:"required"`
}

func GetUserList(ctx *gin.Context) {
	users, err := models.UserList()
	if err != nil {
		ctx.ProtoBuf(400, gin.H{"msg": err})
		return
	}
	var result []*user.User
	for _, item := range users {
		result = append(result, item.ToProtoBuf())
	}
	userList := user.UserList{
		Users: result,
	}
	ctx.ProtoBuf(200, &userList)
}

func GetUserDetail(ctx *gin.Context) {
	var params UserDetailParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}
	u, err := models.FilterByID(params.ID)
	if err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}
	ctx.JSON(200, u.ToGinH())
}
