package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// naming convention: ping is route, get is http method
// function name must be begin with uppercase if we want to call in main func
// func PingGet(c *gin.Context) {
// 	c.JSON(http.StatusOK, map[string]string{
// 		"hello": "Found me",
// 	})
// }

// Better approach so we can import dependecy to PingGet function
func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "Found me",
		})
	}
}
