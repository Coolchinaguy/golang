package hander

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"index.html",gin.H{
		"title" : "hello gin " +
			strings.ToLower(ctx.Request.Method) + " method",
	})
	ctx.HTML(http.StatusOK,"header.html",gin.H{

	})
}
