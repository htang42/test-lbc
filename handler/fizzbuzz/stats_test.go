package fizzbuzz

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/htang42/test-lbc/middleware/counter"
)

func TestStats(t *testing.T) {
	tests := []struct {
		name     string
		requests []FizzbuzzRequest
		want     []FizzbuzzStatsResponse
	}{
		{
			name:     "no request was send",
			requests: []FizzbuzzRequest{},
			want:     []FizzbuzzStatsResponse{},
		},
		{
			name: "one request was send",
			requests: []FizzbuzzRequest{
				{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"},
			},
			want: []FizzbuzzStatsResponse{
				{FizzbuzzRequest: &FizzbuzzRequest{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"}, Count: 1},
			},
		},
		{
			name: "many requests were send",
			requests: []FizzbuzzRequest{
				{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"},
				{Int1: 2, Int2: 7, Limit: 15, Str1: "salut", Str2: "toto"},
			},
			want: []FizzbuzzStatsResponse{
				{FizzbuzzRequest: &FizzbuzzRequest{Int1: 2, Int2: 7, Limit: 15, Str1: "salut", Str2: "toto"}, Count: 1},
				{FizzbuzzRequest: &FizzbuzzRequest{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"}, Count: 1},
			},
		},
		{
			name: "2nd test many were requests send",
			requests: []FizzbuzzRequest{
				{Int1: 4, Int2: 4, Limit: 20, Str1: "same", Str2: "meme"},
				{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"},
				{Int1: 2, Int2: 7, Limit: 15, Str1: "salut", Str2: "toto"},
				{Int1: 4, Int2: 4, Limit: 20, Str1: "same", Str2: "meme"},
				{Int1: 4, Int2: 4, Limit: 20, Str1: "same", Str2: "meme"},
				{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"},
				{Int1: 2, Int2: 7, Limit: 15, Str1: "salut", Str2: "toto"},
				{Int1: 2, Int2: 7, Limit: 15, Str1: "salut", Str2: "toto"},
			},
			want: []FizzbuzzStatsResponse{
				{FizzbuzzRequest: &FizzbuzzRequest{Int1: 2, Int2: 7, Limit: 15, Str1: "salut", Str2: "toto"}, Count: 3},
				{FizzbuzzRequest: &FizzbuzzRequest{Int1: 4, Int2: 4, Limit: 20, Str1: "same", Str2: "meme"}, Count: 3},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := gin.Default()
			h := &FizzbuzzHandler{rc: counter.NewAVLRequestCounter(), limit: 100}

			fizzbuzzRouter := r.Group("/fizzbuzz")
			{
				fizzbuzzRouter.GET("", counter.Counter(h.rc, h.ConvertFizzbuzzRequestAsInterface), h.Fizzbuzz)
				fizzbuzzRouter.GET("/stats", h.Stats)
			}
			size := len(test.requests)
			for x := 0; x < size; x++ {
				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", fmt.Sprintf("/fizzbuzz?int1=%d&int2=%d&limit=%d&str1=%s&str2=%s", test.requests[x].Int1, test.requests[x].Int2, test.requests[x].Limit, test.requests[x].Str1, test.requests[x].Str2), nil)
				r.ServeHTTP(w, req)

			}
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/fizzbuzz/stats", nil)
			r.ServeHTTP(w, req)

			var got []FizzbuzzStatsResponse
			json.Unmarshal(w.Body.Bytes(), &got)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want: %+v\ngot : %+v\n", test.want, got)
			}
		})
	}
}
