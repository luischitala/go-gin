package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luischitala/golang-gin/entity"
	"github.com/luischitala/golang-gin/service"
)

type VideoController interface {
	Save(ctx *gin.Context) entity.Video
	FindAll() []entity.Video
}

// Struct that will implement video controller interface

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	// Delegate the service
	c.service.Save(video)
	return video
}
