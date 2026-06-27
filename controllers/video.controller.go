package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/seth/golang-gin-poc/entity"
	"github.com/seth/golang-gin-poc/service"
)



type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) 
}


type controller struct {
	service service.VideoService
}


func New(service service.VideoService) VideoController {
	return controller {
		service: service,
	}
}