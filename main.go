package main

import (
	"github.com/gin-gonic/gin"
	ex "github.com/iLLia12/studing-api/pkg/executor"
	"net/http"
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

	router.POST("/execute", runCode)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func runCode(c *gin.Context) {
	// Fetch POST data from Context and bind it to the data var
	var data ex.CodeExecuteData
	if err := c.BindJSON(&data); err != nil {
		return
	}

	output := ex.Run(data)

	//fmt.Println(string(output))

	c.IndentedJSON(http.StatusOK, string(output))
}
