package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpModel "github.com/harisaginting/gwyn/models/http"
	"github.com/harisaginting/gwyn/utils/helper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var page httpModel.Page

func init() {
	page = httpModel.Page{
		Now:    helper.Now().Format("2006-01-02 15:04:05"),
		Domain: helper.MustGetEnv("DOMAIN"),
	}
}

func Page(r *gin.RouterGroup) {
	// routing page
	r.GET("/", homepage)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"p": page,
	})
}
