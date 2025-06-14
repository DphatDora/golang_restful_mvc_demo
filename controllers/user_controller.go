package controllers

import (
	"go_restful_mvc/config"
	"go_restful_mvc/dto/req"
	"go_restful_mvc/dto/res"
	"go_restful_mvc/models"
	"go_restful_mvc/services"
	"go_restful_mvc/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (ctrl *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.service.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Register failed"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) Login(c *gin.Context) {
	var request req.LoginRequest
	var response res.LoginResponse

	// Get login attempts from Redis
	attempts, key := utils.GetLoginAttempts(c)

	if attempts >= 5 {
		response.Message = "Too many login attempts. Please try again later."
		response.Token = ""
		c.JSON(http.StatusTooManyRequests, response)
		return
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := ctrl.service.Login(request.Email, request.Password)
	if err != nil {
		// Increment login attempts in Redis
		config.RedisClient.Incr(c, key)
		config.RedisClient.Expire(c, key, 1*time.Minute) // Set expiration for 1 minutes

		response.Message = "Login failed"
		response.Token = ""
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		response.Message = "Error generating token"
		response.Token = ""
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Reset login attempts on successful login
	config.RedisClient.Del(c, key)
	response.Message = "Login successful"
	response.Token = token
	c.JSON(http.StatusOK, response)
}

func (ctrl *UserController) Update(c *gin.Context) {
	id := c.Param("id")

	user, err := ctrl.service.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Name = c.PostForm("name")

	if imagePath, exists := c.Get("image_url"); exists {
		user.ImageURL = imagePath.(string)
	}

	if err := ctrl.service.Update(id, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated Successfully"})

}
