package fizzbuzz

import "github.com/gin-gonic/gin"

type FizzbuzzHandler struct {
}

func RegisterRoutes(r *gin.Engine) {
	h := &FizzbuzzHandler{}

	fizzbuzzRouter := r.Group("/fizzbuzz")
	{
		fizzbuzzRouter.GET("", h.Fizzbuzz)
	}
}
