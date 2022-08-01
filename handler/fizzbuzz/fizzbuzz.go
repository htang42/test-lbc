package fizzbuzz

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FizzbuzzRequest struct {
	Int1  int    `form:"int1" binding:"required,gte=1,lte=100"`
	Int2  int    `form:"int2" binding:"required,gte=1,lte=100"`
	Limit int    `form:"limit" binding:"required,gte=1,lte=100"`
	Str1  string `form:"str1" binding:"required"`
	Str2  string `form:"str2" binding:"required"`
}

// needed to be able to use the AVL request counter middleware
func (fbr FizzbuzzRequest) GetKey() string {
	return fmt.Sprintf("%d_%d_%d_%s_%s", fbr.Int1, fbr.Int2, fbr.Limit, fbr.Str1, fbr.Str2)
}

// need to be able to use the counter middleware
func (h *FizzbuzzHandler) ConvertFizzbuzzRequestAsInterface(c *gin.Context) (interface{}, error) {
	return h.GetFizzbuzzRequest(c)
}

func (h *FizzbuzzHandler) GetFizzbuzzRequest(c *gin.Context) (*FizzbuzzRequest, error) {
	var req FizzbuzzRequest
	if err := c.ShouldBind(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func (h *FizzbuzzHandler) Fizzbuzz(c *gin.Context) {
	req, err := h.GetFizzbuzzRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := fizzbuzz(req.Int1, req.Int2, req.Limit, req.Str1, req.Str2)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func fizzbuzz(int1, int2, limit int, str1, str2 string) ([]string, error) {
	if int1 < 1 || int1 > 100 {
		return nil, fmt.Errorf("int1 should be between 1 and 100")
	}
	if int2 < 1 || int2 > 100 {
		return nil, fmt.Errorf("int2 should be between 1 and 100")
	}
	if limit < 1 || limit > 100 {
		return nil, fmt.Errorf("limit should be between 1 and 100")
	}
	result := make([]string, limit)
	for x := 1; x <= limit; x++ {
		str := ""
		if x%int1 == 0 {
			str = str1
		}
		if x%int2 == 0 {
			str += str2
		}
		if len(str) == 0 {
			str = strconv.Itoa(x)
		}
		result[x-1] = str
	}
	return result, nil
}
