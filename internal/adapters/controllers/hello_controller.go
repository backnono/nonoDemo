package controllers

import (
	"github.com/gin-gonic/gin"
	"nonoDemo/pkg/adapters/agin"
	"nonoDemo/pkg/framework"
)

type HelloController struct {
	logger framework.Logger
}

func NewHelloController(logger framework.Logger) *HelloController {
	return &HelloController{
		logger: logger,
	}
}

func (ctrl *HelloController) InitRouter(r *gin.Engine) {
	r.GET("/api/hello/v1", ctrl.get)
	r.GET("/api/hello/v1/err", ctrl.err)
}

func (ctrl *HelloController) get(c *gin.Context) {
	agin.WriteResponse(c, agin.ResponseData{Data: map[string]string{"msg": "hello world!"}}, nil)
}

func (ctrl *HelloController) err(c *gin.Context) {
	err := agin.NewHTTPError(framework.CommonInternalErr.New("common error").(framework.AppError), 500)
	agin.WriteResponse(c, agin.ResponseData{Data: map[string]string{"msg": "hello world!"}}, err)
}
