package fizzbuzz

import (
	"github.com/gin-gonic/gin"
	"github.com/htang42/test-lbc/handler/middleware/counter"
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
	}
}
