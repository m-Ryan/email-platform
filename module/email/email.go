package email

import "github.com/gin-gonic/gin"

func Module(path string, router *gin.Engine) {
	controller(router.Group(path))
}
