package fizzbuzz

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/htang42/test-lbc/handler/middleware/counter"
)

type FizzbuzzStatsResponse struct {
	*FizzbuzzRequest
	Count int
}

func (h *FizzbuzzHandler) Stats(c *gin.Context) {
	mcr, err := h.rc.FindMostCalledRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ar := mcr.([]*counter.AVLNodeCounter)
	res := make([]*FizzbuzzStatsResponse, len(ar))
	size := len(ar)
	for x := 0; x < size; x++ {
		res[x] = &FizzbuzzStatsResponse{
			FizzbuzzRequest: ar[x].AVLNodeData.(*FizzbuzzRequest),
			Count:           ar[x].Count,
		}
	}
	c.JSON(http.StatusOK, res)
}
