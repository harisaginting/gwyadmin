package response

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/harisaginting/gwyn/utils/helper"
)

type Message struct {
	Message string `json:"message"`
}

func DefaultAppHeader(c *gin.Context) {
	tm := time.Now().Unix()
	c.Writer.Header().Set("App-Name", helper.MustGetEnv("APP_NAME"))
	c.Writer.Header().Set("App-Version", helper.MustGetEnv("APP_VERSION"))
	c.Writer.Header().Set("App-Time", strconv.Itoa(int(tm)))
}

func Json(c *gin.Context, data interface{}) {
	DefaultAppHeader(c)
	c.JSON(http.StatusOK, gin.H{
		"status":        http.StatusOK,
		"data":          data,
		"error_message": nil,
	})
	return
}

func StatusOK(c *gin.Context, data interface{}) {
	if data != nil {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusOK, nil)
	}

}
func StatusNotFound(c *gin.Context, err error) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusNotFound, gin.H{
		"message": err.Error(),
	})

}

func StatusRedirect(c *gin.Context, url string) {
	c.Writer.Header().Set("Location", url)
	c.Redirect(http.StatusFound, url)

}

func StatusCreated(c *gin.Context, data interface{}) {
	if data != nil {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(http.StatusCreated, data)
	} else {
		c.JSON(http.StatusCreated, nil)
	}

}

func StatusConflict(c *gin.Context, err error) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusConflict, gin.H{
		"message": err.Error(),
	})
}

func StatusUnprocessableEntity(c *gin.Context, err error) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": err.Error(),
	})
}

func StatusError(c *gin.Context, err error) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
	})
}

func BadRequest(c *gin.Context, s ...string) {
	remark := "The request is not valid in this context"
	if len(s) > 0 {
		if s[0] != "" {
			remark = s[0]
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": remark,
	})
}

func NoContent(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNoContent)
}

func Accepted(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusAccepted)
}
