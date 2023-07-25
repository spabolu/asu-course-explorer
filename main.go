package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var client *redis.Client

// Redis cache layer
// Check if a value exists for the given key in cache
func cache(key string) (string, error) {
	return client.Get(ctx, key).Result()
}

// Function to store a key-value pair in the cache
func setCache(key string, value string) error {
	return client.Set(ctx, key, value, 30*time.Minute).Err()
}

// Handler function to handle the '/api/classes' endpoint
// It checks the cache before making a scraping operation
func getClasses(c *gin.Context) {
	cached, err := cache("classes")
	if err == nil && cached != "" {
		var classes []ClassInformation
		if err := json.Unmarshal([]byte(cached), &classes); err == nil {
			c.IndentedJSON(http.StatusOK, classes)
			fmt.Println("Cache hit")
			return
		}
	}

	fmt.Println("Cache miss")
	classes := scrapeClasses()
	data, err := json.Marshal(classes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshalling class data"})
		return
	}

	setCache("classes", string(data))
	c.IndentedJSON(http.StatusOK, classes)
}

// Handler function to handle the '/api/classes/:classId' endpoint
// It checks the cache before making a scraping operation
func getClassById(c *gin.Context) {
	classId := c.Param("classId")

	cached, err := cache(classId)
	if err == nil && cached != "" {
		var class []ClassInformation
		if err := json.Unmarshal([]byte(cached), &class); err == nil {
			c.IndentedJSON(http.StatusOK, class)
			fmt.Println("Cache hit")
			return
		}
	}

	fmt.Println("Cache miss")
	class := scrapeClass(classId)
	data, err := json.Marshal(class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshalling class data"})
		return
	}

	setCache(classId, string(data))
	c.IndentedJSON(http.StatusOK, class)
}

// Handler function to handle the '/api/classes/course/:courseId' endpoint
// It checks the cache before making a scraping operation
func getClassesByCourse(c *gin.Context) {
	courseId := c.Param("courseId")

	cached, err := cache(courseId)
	if err == nil && cached != "[]" && cached != "" {
		var classes []ClassInformation
		if err := json.Unmarshal([]byte(cached), &classes); err == nil {
			c.IndentedJSON(http.StatusOK, classes)
			fmt.Println("Cache hit")
			return
		}
	}

	fmt.Println("Cache miss")
	subject := courseId[:3]
	number := courseId[3:]

	classes := scrapeClassesByCourse(subject, number)
	data, err := json.Marshal(classes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshalling class data"})
		return
	}

	setCache(courseId, string(data))
	c.IndentedJSON(http.StatusOK, classes)
}

func main() {
	opt, _ := redis.ParseURL("redis://default:c935507737d5440e9c3ea40001e832cd@usw1-equal-crow-33656.upstash.io:33656")
	client = redis.NewClient(opt)

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
