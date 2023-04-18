package agin

import "github.com/gin-gonic/gin"

type Controller interface {
	InitRouter(r *gin.Engine)
}
