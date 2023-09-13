package api

import (
	"github.com/gin-gonic/gin"
	"github.com/harisaginting/gwyn/routers/api/shorten"
)

func Implement(group *gin.RouterGroup) {
	group = group.Group("api")
	//implement group shorten
	shorten.Add(group)
}
