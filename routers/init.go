package routers

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	e := gin.Default()
	e.Static("/html", "./html")
	api := e.Group("/api")
	api.POST("/registration", account.CreateAccount)
	api.POST("/login", account.Authorization)
	account.Account(e.Group("/api/user", CheckToken))
	return e
}
