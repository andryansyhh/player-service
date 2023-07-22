package handler

import (
	"player-service/src/usecase"
)

type handlerHttpServer struct {
	usecase usecase.UsecaseInterface
}

func NewHttpHandler(usecase usecase.UsecaseInterface) handlerHttpServer {
	return handlerHttpServer{usecase: usecase}
}
