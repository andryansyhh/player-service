package handler

import (
	"net/http"
	"player-service/src/models"

	"github.com/gin-gonic/gin"
)

func (h *handlerHttpServer) CreateUser(c *gin.Context) {
	var req models.CreateUser
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	err = h.usecase.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "success create user")
}

func (h *handlerHttpServer) Login(c *gin.Context) {
	var input models.Login
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	res, err := h.usecase.Login(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handlerHttpServer) Logout(c *gin.Context) {
	var input models.Logout
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	err = h.usecase.Logout(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "success logout")
}

func (h *handlerHttpServer) GetUserByUuid(c *gin.Context) {
	userUuid := c.Param("uuid")
	res, err := h.usecase.GetUserByUuid(userUuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handlerHttpServer) GetAllUsers(c *gin.Context) {
	var req models.ListRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	res, err := h.usecase.GetAllUsers(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
