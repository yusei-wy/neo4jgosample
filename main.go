package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const PORT = 8080

func main() {
	router := gin.Default()

	router.GET("/", handleHome)

	router.Run(fmt.Sprintf(":%d", PORT))
}

func handleHome(c *gin.Context) {
	c.String(http.StatusOK, "HELLO WORLD")
}
