package controller

import (
	"goapp-gin/entity"
	"goapp-gin/service"
	"goapp-gin/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	FindAll() []entity.Video
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service1 service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("isGoodTitle", validators.ValidateVideoTitle)
	return &controller{
		service: service1,
	}
}

func (con *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	con.service.Save(video)
	return nil
}

func (con *controller) Update(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	con.service.Update(video)
	return nil
}

func (con *controller) Delete(ctx *gin.Context) error {
	var video entity.Video
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id
	con.service.Delete(video)
	return nil
}

func (con *controller) FindAll() []entity.Video {
	return con.service.FindAll()
}

func (con *controller) ShowAll(ctx *gin.Context) {
	videos := con.service.FindAll()
	data := gin.H{
		"title":  "Videos Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
