package router

import (
	"github.com/gin-gonic/gin"
	"mxaof_blog/handler/login"
)

func NewRouter() {
	router := gin.New()
	loginRouter := router.Group("/login")
	loginRouter.GET("/captcha", login.GetCaptchaCode)
	router.Run(":80")
}
