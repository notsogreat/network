package main

import (
	"net/http"
	"os/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize a Gin router
	r := gin.Default()

	// Use the CORS middleware with default configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://example.com", "http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Authorization"}

	r.Use(cors.New(config))

	// Define a route handler for the API
	api := r.Group("/api")
	{
		api.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello, world!",
			})
		})

		api.GET("/info", func(c *gin.Context) {
			cmd := exec.Command("bash", "run-script.sh")

			output, err := cmd.CombinedOutput()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Failed to run script",
					"message": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Script executed successfully",
				"output":  string(output),
			})
		})
	}

	// Serve the frontend build
	r.StaticFile("/", "../frontend/build/index.html")
	r.Static("/static", "../frontend/build/static")

	// Run the server on port 8080
	r.Run(":8080")
}
