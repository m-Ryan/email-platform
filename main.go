package main

import (
	"github.com/gin-gonic/gin"
	"github.com/m-ryan/tool-platform/module/email"
)

func main() {
	router := gin.Default()

	email.Module("/email", router)
	router.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
