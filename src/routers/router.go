package routers

import (
	"player-service/src/handler"
	"player-service/src/middleware"
	"player-service/src/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app usecase.UsecaseInterface) {
	handler := handler.NewHttpHandler(app)
	user := r.Group("/user")
	{
		user.POST("/create-user", handler.CreateUser)
		user.POST("/login", handler.Login)
		user.POST("/logout", handler.Logout)
		user.Use(middleware.TokenCheckerMiddleware)
		user.GET("/get/:uuid", handler.GetUserByUuid)
		user.POST("/all-user", handler.GetAllUsers)
	}

	userWallet := r.Group("/wallet")
	{
		userWallet.Use(middleware.TokenCheckerMiddleware)
		userWallet.POST("/create-user-wallet", handler.CreateUserWallet)
		userWallet.POST("/topup", handler.TopupUserWallet)
	}

}
