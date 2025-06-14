package utils

import (
	"fmt"
	"go_restful_mvc/config"

	"github.com/gin-gonic/gin"
)

func GetClientIP(c *gin.Context) string {
	ip := c.ClientIP()
	if ip == "" {
		ip = "unknown"
	}
	return ip
}

func GetLoginAttempts(c *gin.Context) (int, string) {
	clientIP := GetClientIP(c)
	key := "login_attempts:" + clientIP
	attemptsStr, _ := config.RedisClient.Get(c, key).Result()
	attempts := 0
	if attemptsStr != "" {
		fmt.Sscanf(attemptsStr, "%d", &attempts)
	}
	return attempts, key
}
