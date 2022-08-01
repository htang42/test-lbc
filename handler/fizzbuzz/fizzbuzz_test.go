package fizzbuzz

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
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
	h := &FizzbuzzHandler{}
	r.GET("/fizzbuzz", h.Fizzbuzz)

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
				t.Errorf("should have a code status bad request %+v\n", w.Code)
			}

		})
	}

}
