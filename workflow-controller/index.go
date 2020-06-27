package controller

import (
	"fmt"
	"go-admin-demo/tools/app"

	"github.com/gin-gonic/gin"
)

// Index 首页
func Index(c *gin.Context) {
	app.OK(c, "Hello world!", "")
}

// GetToken 获取token
func GetToken(c *gin.Context) (string, error) {
	token, exists := c.Get("JWT_TOKEN")
	if !exists {
		return "", fmt.Errorf("JWT_TOKEN NOT EXISTED")
	}

	return token.(string), nil
}
