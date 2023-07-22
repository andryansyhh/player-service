package handler

import (
	"net/http"
	"player-service/src/models"

	"github.com/gin-gonic/gin"
)

func (h *handlerHttpServer) CreateUserWallet(c *gin.Context) {
	userUuid := c.GetString("user_uuid")
	var req models.CreateUserWallet
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	req.UserUuid = userUuid

	err = h.usecase.CreateUserWallet(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "success create user wallet")
}

func (h *handlerHttpServer) TopupUserWallet(c *gin.Context) {
	userUuid := c.GetString("user_uuid")

	var req models.TopupUserWallet
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	err = h.usecase.TopupUserWallet(c, req, userUuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "success topup wallet")
}
