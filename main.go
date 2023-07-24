package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getClassData(c *gin.Context) {
	classNumber := c.Param("classNumber")
	ClassInformation, err := scrapeClass(classNumber)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, ClassInformation)
}

func main() {
	router := gin.Default()
	router.GET("/class/:classNumber", getClassData)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	} else {
		fmt.Println("Server is running on port 8080")
	}
}
