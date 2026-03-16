package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ex "github.com/iLLia12/studing-api/pkg/runner"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/hello", test)
	router.POST("/execute", runCode)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello world",})
}

func runCode(c *gin.Context) {
	var data ex.Payload
	if err := c.BindJSON(&data); err != nil { return }
	output := ex.Run(data)
	c.JSON(http.StatusOK, gin.H{
		"output": string(replaceNewlineWithCRLF(output)),
	})
}

func replaceNewlineWithCRLF(input []byte) []byte {
	var output []byte
	for _, b := range input {
		if b == 10 {
			output = append(output, 13, 10) // Add CRLF
		} else {
			output = append(output, b) // Add the original byte
		}
	}
	return output
}
