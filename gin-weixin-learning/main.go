package main

import "gin-weixin-learning/initRouter"

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}