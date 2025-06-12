package controllers

import (
	"go_restful_mvc/dto/req"
	"go_restful_mvc/dto/res"
	"go_restful_mvc/models"
	"go_restful_mvc/services"
	"go_restful_mvc/utils"
	"net/http"

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

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := ctrl.service.Login(request.Email, request.Password)
	if err != nil {
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

	response.Message = "Login successful"
	response.Token = token
	c.JSON(http.StatusOK, response)
}

func (ctrl *UserController) Update(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.service.Update(id, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}
