package agin

import (
	"nonoDemo/pkg/framework"
	"nonoDemo/pkg/utils/observability"
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    framework.ErrorType `json:"code"`
	Message string              `json:"message"`
	Data    any                 `json:"data"`
}

func WriteResponse(c *gin.Context, response ResponseData, err error) {
	if err != nil {
		if httpErr, ok := err.(HTTPError); ok {
			c.JSON(httpErr.code, ResponseData{
				Code:    framework.GetType(err),
				Message: httpErr.AppError.Error(),
			})
			observability.AppendEvents(c.Request.Context(),
				"biz.code", framework.GetType(err),
				"err.msg", httpErr.AppError.Error(),
			)
		} else if appErr, ok := err.(framework.AppError); ok {
			c.JSON(500, ResponseData{
				Code:    framework.GetType(err),
				Message: appErr.Error(),
			})
			observability.AppendEvents(c.Request.Context(),
				"biz.code", framework.GetType(err),
				"err.msg", appErr.Error(),
			)
		} else {
			c.JSON(500, ResponseData{
				Code:    framework.CommonInternalErr,
				Message: err.Error(),
			})
			observability.AppendEvents(c.Request.Context(),
				"biz.code", framework.CommonInternalErr,
				"err.msg", err.Error(),
			)
		}
	} else {
		c.JSON(200, response)
	}
}
