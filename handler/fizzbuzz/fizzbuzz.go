package fizzbuzz

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FizzbuzzRequest struct {
	Int1  int    `form:"int1" binding:"required,gte=1"`
	Int2  int    `form:"int2" binding:"required,gte=1"`
	Limit int    `form:"limit" binding:"required,gte=1"`
	Str1  string `form:"str1" binding:"required"`
	Str2  string `form:"str2" binding:"required"`
}

// needed to be able to use the AVL request counter middleware
// it returns a unique string by concatenating the fields with _
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
	if req.Int1 > h.limit {
		return nil, fmt.Errorf("int1 should be less than or equal to %d", h.limit)
	}
	if req.Int2 > h.limit {
		return nil, fmt.Errorf("int2 should be less than or equal to %d", h.limit)
	}
	if req.Limit > h.limit {
		return nil, fmt.Errorf("limit should be less than or equal to %d", h.limit)
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
	if h.rs != nil {
		h.rs.SetResponse(req, result)
	}
	c.JSON(http.StatusOK, result)
}

func fizzbuzz(int1, int2, limit int, str1, str2 string) ([]string, error) {
	if int1 < 1 {
		return nil, fmt.Errorf("int1 should be greater than 0")
	}
	if int2 < 1 {
		return nil, fmt.Errorf("int2 should be greater than 0")
	}
	if limit < 1 {
		return nil, fmt.Errorf("limit should be greater than 0")
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
