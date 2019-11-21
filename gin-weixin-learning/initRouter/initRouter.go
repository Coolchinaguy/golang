package initRouter

import (
	"gin-weixin-learning/hander"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/",hander.Index)
	router.GET("/login",)
	userRouter := router.Group("/user")
	{
			userRouter.GET("/:name",hander.Save)
			userRouter.GET("",hander.UserSaveByquery)
	}
	router.LoadHTMLGlob("gin-weixin-learning/templates/*")
	return router
}




/*func retHelloGinAndMethod(context *gin.Context){
	context.String(http.StatusOK,
		"hello gin" + " " + strings.ToLower(context.Request.Method) + " " + "method")
}*/
