package fizzbuzz

import (
	"github.com/gin-gonic/gin"
	"github.com/htang42/test-lbc/middleware/counter"
	responseSaver "github.com/htang42/test-lbc/middleware/response-saver"
)

type FizzbuzzHandler struct {
	rc    counter.RequestCounter
	rs    responseSaver.ResponseSaver
	limit int
}

func RegisterRoutes(r *gin.Engine, limit int) {
	h := &FizzbuzzHandler{
		limit: limit,
		rc:    counter.NewAVLRequestCounter(),
		rs:    responseSaver.NewAVLResponseSaver(),
	}

	fizzbuzzRouter := r.Group("/fizzbuzz")
	{
		fizzbuzzRouter.GET("",
			counter.Counter(h.rc, h.ConvertFizzbuzzRequestAsInterface),
			responseSaver.Retrieve(h.rs, h.ConvertFizzbuzzRequestAsInterface),
			h.Fizzbuzz,
		)
		fizzbuzzRouter.GET("/stats", h.Stats)
	}
}
