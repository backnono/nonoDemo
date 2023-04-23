// Package controllers @Author nono.he 2023/4/20 16:00:00
package controllers

import (
	"github.com/gin-gonic/gin"

	"nonoDemo/internal/application/executor"
	"nonoDemo/pkg/adapters/agin"
	"nonoDemo/pkg/framework"
)

type BlogController struct {
	logger   framework.Logger
	blogExec *executor.BlogOperator
}

func NewBlogController(logger framework.Logger, blogExec *executor.BlogOperator) *BlogController {
	return &BlogController{
		logger:   logger,
		blogExec: blogExec,
	}
}

func (ctrl *BlogController) InitRouter(r *gin.Engine) {
	r.GET("/api/blog/v1/:blog_id", ctrl.getBlog)
	//r.GET("/api/blog/v1/err", ctrl.err)
}

func (ctrl *BlogController) getBlog(c *gin.Context) {
	blogID := c.Param("blog_id")
	blog, err := ctrl.blogExec.GetBlog(c.Request.Context(), blogID)
	if err != nil {
		err = agin.NewHTTPError(err, 500)
	}
	agin.WriteResponse(c, agin.ResponseData{Data: blog}, err)
}
