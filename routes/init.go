package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ilya-sokolov/crypto_kiddies-server/routes/user"
)

func Router() *gin.Engine {
	e := gin.Default()
	api := e.Group("/api")
	api.POST("/registration", user.CreateAccount)
	api.POST("/login", user.Authorization)
	user.Account(e.Group("/api/user", CheckToken))
	return e
}
