package hander

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Save(ctx *gin.Context) {
	username := ctx.Param("name")
	ctx.String(http.StatusOK,username + " " + "用户已经保存")
}

func UserSaveByquery(ctx *gin.Context) {
	username := ctx.Query("name")
	age := ctx.Query("age")
	ctx.String(http.StatusOK,
		"用户： " + username + ", 年龄： " + age + " 已经保存")
}