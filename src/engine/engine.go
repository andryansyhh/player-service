package engine

import (
	"fmt"
	"log"
	"os"

	"player-service/src/config"
	"player-service/src/models"
	"player-service/src/repository"
	"player-service/src/routers"
	"player-service/src/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StartApplication() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	db := config.GetDatabase(os.Getenv("database"))
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.UserWallet{})

	r := gin.New()
	Repo := repository.NewRepo(db)
	app := usecase.NewUsecase(Repo)
	routers.RegisterApi(r, app)

	port := os.Getenv("APP_PORT")
	r.Run(fmt.Sprintf(":%s", port))
}
