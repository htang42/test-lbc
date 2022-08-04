package counter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestCounter interface {
	IncrementCounter(interface{}) error
	FindMostCalledRequests() (interface{}, error)
}

func Counter(rc RequestCounter, getRequest func(c *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := getRequest(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := rc.IncrementCounter(req); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}
