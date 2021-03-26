package email

import "github.com/gin-gonic/gin"

func controller(r *gin.RouterGroup) {
	r.GET("/", getEmail)
	r.POST("/", sendEnail)
}
