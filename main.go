package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler function to handle the '/api/classes' endpoint
// It sends back JSON data of the first ~75 classes
func getClasses(c *gin.Context) {
	classes := scrapeClasses()
	c.IndentedJSON(http.StatusOK, classes)
}

// Handler function to handle the '/api/classes/:classId' endpoint
// It sends back JSON data of a class with a specific ID
func getClassById(c *gin.Context) {
	classId := c.Param("classId")
	class := scrapeClass(classId)
	c.IndentedJSON(http.StatusOK, class)
}

// Handler function to handle the '/api/classes/course/:courseId' endpoint
// It sends back JSON data of a class with a specific ID
func getClassesByCourse(c *gin.Context) {
	subject := c.Param("courseId")[:3]
	number := c.Param("courseId")[3:]

	classes := scrapeClassesByCourse(subject, number)
	c.IndentedJSON(http.StatusOK, classes)
}

func main() {
	router := gin.Default()

	router.GET("/api/classes", getClasses)
	router.GET("/api/classes/:classId", getClassById)
	router.GET("/api/classes/course/:courseId", getClassesByCourse)

	// Run the server and handle possible errors
	if err := router.Run(); err != nil {
		fmt.Printf("Server failed to run: %v", err)
	} else {
		fmt.Println("Server is running on port 8080")
	}
}
