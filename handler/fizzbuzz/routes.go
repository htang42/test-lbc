package fizzbuzz

import (
	"github.com/gin-gonic/gin"
	"github.com/htang42/test-lbc/middleware/counter"
)

type FizzbuzzHandler struct {
	rc counter.RequestCounter
}

func RegisterRoutes(r *gin.Engine) {
	h := &FizzbuzzHandler{
		rc: counter.NewAVLRequestCounter(),
	}

	fizzbuzzRouter := r.Group("/fizzbuzz")
	{
		fizzbuzzRouter.GET("", counter.Counter(h.rc, h.ConvertFizzbuzzRequestAsInterface), h.Fizzbuzz)
		fizzbuzzRouter.GET("/stats", h.Stats)
	}
}
