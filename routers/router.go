package router

import (
	"github.com/gin-gonic/gin"
	api "github.com/harisaginting/gwyn/routers/api"
)

// Swagger Config
// @title gwyn
// @version 1.0
// @description gwyn
// @host localhost:4000
// @BasePath /
// @schemes http
// @query.collection.format multi
// @contact.name Harisa Ginting
// @contact.url ”
func Api(r *gin.RouterGroup) {
	// api
	api.Implement(r)
}
