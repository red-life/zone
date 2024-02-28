package http

import "github.com/gin-gonic/gin"

func BasicAuth(username string, password string) gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		username: password,
	})
}
