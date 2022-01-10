package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", rootHandler)

	r.GET("/product/:id", productHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "root",
	})
}

func productHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}
