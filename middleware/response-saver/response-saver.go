package responseSaver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseSaver interface {
	GetResponse(interface{}) (interface{}, error)
	// need to be called by the handler that compute the response
	// otherwise the response will not be saved
	SetResponse(interface{}, interface{}) error
}

func Retrieve(rs ResponseSaver, getRequest func(c *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := getRequest(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, err := rs.GetResponse(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if result != nil {
			c.AbortWithStatusJSON(http.StatusOK, result)
			return
		}
	}
}
