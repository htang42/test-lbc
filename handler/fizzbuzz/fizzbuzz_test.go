package fizzbuzz

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	responseSaver "github.com/htang42/test-lbc/middleware/response-saver"
)

func TestFizzbuzz(t *testing.T) {
	failedTests := []struct {
		name  string
		int1  int
		int2  int
		limit int
		str1  string
		str2  string
	}{
		{name: "int1 is < 1", int1: -1, int2: 5, limit: 100, str1: "fizz", str2: "buzz"},
		{name: "int1 is > 100", int1: 101, int2: 5, limit: 100, str1: "fizz", str2: "buzz"},
		{name: "int2 is < 1", int1: 3, int2: -1, limit: 100, str1: "fizz", str2: "buzz"},
		{name: "int2 is > 100", int1: 3, int2: 101, limit: 100, str1: "fizz", str2: "buzz"},
		{name: "limit is < 1", int1: 3, int2: 5, limit: 0, str1: "fizz", str2: "buzz"},
		{name: "limit is > 100", int1: 3, int2: 5, limit: 1000, str1: "fizz", str2: "buzz"},
	}
	for _, test := range failedTests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := fizzbuzz(test.int1, test.int2, test.limit, test.str1, test.str2); err == nil {
				t.Errorf("should get an error")
			}
		})
	}

	successTests := []struct {
		name  string
		int1  int
		int2  int
		limit int
		str1  string
		str2  string
		want  []string
	}{
		{
			name: "first success test", int1: 3, int2: 5, limit: 100, str1: "Fizz", str2: "Buzz",
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14",
				"FizzBuzz", "16", "17", "Fizz", "19", "Buzz", "Fizz", "22", "23", "Fizz", "Buzz", "26",
				"Fizz", "28", "29", "FizzBuzz", "31", "32", "Fizz", "34", "Buzz", "Fizz", "37", "38",
				"Fizz", "Buzz", "41", "Fizz", "43", "44", "FizzBuzz", "46", "47", "Fizz", "49", "Buzz",
				"Fizz", "52", "53", "Fizz", "Buzz", "56", "Fizz", "58", "59", "FizzBuzz", "61", "62",
				"Fizz", "64", "Buzz", "Fizz", "67", "68", "Fizz", "Buzz", "71", "Fizz", "73", "74",
				"FizzBuzz", "76", "77", "Fizz", "79", "Buzz", "Fizz", "82", "83", "Fizz", "Buzz", "86",
				"Fizz", "88", "89", "FizzBuzz", "91", "92", "Fizz", "94", "Buzz", "Fizz", "97", "98",
				"Fizz", "Buzz"},
		},
		{
			name: "second success test", int1: 2, int2: 4, limit: 10, str1: "sa", str2: "lut",
			want: []string{"1", "sa", "3", "salut", "5", "sa", "7", "salut", "9", "sa"},
		},
		{
			name: "third success test", int1: 3, int2: 3, limit: 14, str1: "les", str2: "Copains",
			want: []string{"1", "2", "lesCopains", "4", "5", "lesCopains", "7", "8", "lesCopains", "10", "11", "lesCopains", "13", "14"},
		},
	}
	for _, test := range successTests {
		t.Run(test.name, func(t *testing.T) {
			got, err := fizzbuzz(test.int1, test.int2, test.limit, test.str1, test.str2)
			if err != nil {
				t.Errorf("should not get an error")
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want: %+v\ngot: %+v\n", test.want, got)
			}
		})
	}
}

func TestHandlerFizzBuzz(t *testing.T) {
	r := gin.Default()
	h := &FizzbuzzHandler{
		rs: responseSaver.NewAVLResponseSaver(),
	}
	r.GET("/fizzbuzz",
		responseSaver.Retrieve(h.rs, h.ConvertFizzbuzzRequestAsInterface),
		h.Fizzbuzz,
	)

	failedTests := []struct {
		name    string
		request string
	}{
		{name: "empty", request: ""},
		{name: "missing int1", request: "int2=5&limit=100&str1=fizz&str2=buzz"},
		{name: "missing int2", request: "int1=3&limit=100&str1=fizz&str2=buzz"},
		{name: "missing limit", request: "int1=3&int2=5&str1=fizz&str2=buzz"},
		{name: "missing str1", request: "int1=3&int2=5limit=100&str2=buzz"},
		{name: "missing str2", request: "int1=3&int2=5limit=100&str1=fizz"},
	}

	for _, test := range failedTests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/fizzbuzz?"+test.request, nil)
			r.ServeHTTP(w, req)
			if w.Code != http.StatusBadRequest {
				t.Errorf("should have a code status bad request, got: %+v\n", w.Code)
			}

		})
	}

	successTests := []struct {
		name    string
		request string
		want    []string
	}{
		{
			name: "first request", request: "int1=3&int2=5&limit=100&str1=fizz&str2=buzz",
			want: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz", "fizz", "22", "23", "fizz", "buzz", "26", "fizz", "28", "29", "fizzbuzz", "31", "32", "fizz", "34", "buzz", "fizz", "37", "38", "fizz", "buzz", "41", "fizz", "43", "44", "fizzbuzz", "46", "47", "fizz", "49", "buzz", "fizz", "52", "53", "fizz", "buzz", "56", "fizz", "58", "59", "fizzbuzz", "61", "62", "fizz", "64", "buzz", "fizz", "67", "68", "fizz", "buzz", "71", "fizz", "73", "74", "fizzbuzz", "76", "77", "fizz", "79", "buzz", "fizz", "82", "83", "fizz", "buzz", "86", "fizz", "88", "89", "fizzbuzz", "91", "92", "fizz", "94", "buzz", "fizz", "97", "98", "fizz", "buzz"},
		},
		{
			name: "second request", request: "int1=3&int2=5&limit=20&str1=fizz&str2=buzz",
			want: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz"},
		},
		{
			name: "third request", request: "int1=3&int2=5&limit=1&str1=fizz&str2=buzz",
			want: []string{"1"},
		},
		{
			name: "fourth request", request: "int1=6&int2=7&limit=35&str1=salut&str2=toto",
			want: []string{"1", "2", "3", "4", "5", "salut", "toto", "8", "9", "10", "11", "salut", "13", "toto", "15", "16", "17", "salut", "19", "20", "toto", "22", "23", "salut", "25", "26", "27", "toto", "29", "salut", "31", "32", "33", "34", "toto"},
		},
		{
			name: "fifth request", request: "int1=1&int2=5&limit=20&str1=fizz&str2=buzz",
			want: []string{"fizz", "fizz", "fizz", "fizz", "fizzbuzz", "fizz", "fizz", "fizz", "fizz", "fizzbuzz", "fizz", "fizz", "fizz", "fizz", "fizzbuzz", "fizz", "fizz", "fizz", "fizz", "fizzbuzz"},
		},
	}

	for _, test := range successTests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/fizzbuzz?"+test.request, nil)
			r.ServeHTTP(w, req)
			if w.Code != http.StatusOK {
				t.Errorf("should have a code status ok, got: %+v\n", w.Code)
			}

			var got []string
			json.Unmarshal(w.Body.Bytes(), &got)

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("want: %v\ngot: %v\n", test.want, got)
			}
		})
	}

	successTests = []struct {
		name    string
		request string
		want    []string
	}{
		{
			name: "first request using only responseSaver middleware", request: "int1=3&int2=5&limit=100&str1=fizz&str2=buzz",
			want: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz", "fizz", "22", "23", "fizz", "buzz", "26", "fizz", "28", "29", "fizzbuzz", "31", "32", "fizz", "34", "buzz", "fizz", "37", "38", "fizz", "buzz", "41", "fizz", "43", "44", "fizzbuzz", "46", "47", "fizz", "49", "buzz", "fizz", "52", "53", "fizz", "buzz", "56", "fizz", "58", "59", "fizzbuzz", "61", "62", "fizz", "64", "buzz", "fizz", "67", "68", "fizz", "buzz", "71", "fizz", "73", "74", "fizzbuzz", "76", "77", "fizz", "79", "buzz", "fizz", "82", "83", "fizz", "buzz", "86", "fizz", "88", "89", "fizzbuzz", "91", "92", "fizz", "94", "buzz", "fizz", "97", "98", "fizz", "buzz"},
		},
		{
			name: "second request only responseSaver middleware", request: "int1=3&int2=5&limit=20&str1=fizz&str2=buzz",
			want: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz"},
		},
		{
			name: "third request only responseSaver middleware", request: "int1=3&int2=5&limit=1&str1=fizz&str2=buzz",
			want: []string{"1"},
		},
		{
			name: "fourth request only responseSaver middleware", request: "int1=6&int2=7&limit=35&str1=salut&str2=toto",
			want: []string{"1", "2", "3", "4", "5", "salut", "toto", "8", "9", "10", "11", "salut", "13", "toto", "15", "16", "17", "salut", "19", "20", "toto", "22", "23", "salut", "25", "26", "27", "toto", "29", "salut", "31", "32", "33", "34", "toto"},
		},
		{
			name: "fifth request only responseSaver middleware", request: "int1=1&int2=5&limit=20&str1=fizz&str2=buzz",
			want: []string{"fizz", "fizz", "fizz", "fizz", "fizzbuzz", "fizz", "fizz", "fizz", "fizz", "fizzbuzz", "fizz", "fizz", "fizz", "fizz", "fizzbuzz", "fizz", "fizz", "fizz", "fizz", "fizzbuzz"},
		},
	}

	r = gin.Default()
	r.GET("/fizzbuzz", responseSaver.Retrieve(h.rs, h.ConvertFizzbuzzRequestAsInterface))
	// we don't need anymore the h.Fizzbuzz, all results should be get by the middleware
	for _, test := range successTests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/fizzbuzz?"+test.request, nil)
			r.ServeHTTP(w, req)
			if w.Code != http.StatusOK {
				t.Errorf("should have a code status ok, got: %+v\n", w.Code)
			}

			var got []string
			json.Unmarshal(w.Body.Bytes(), &got)

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("want: %v\ngot: %v\n", test.want, got)
			}
		})
	}
}
