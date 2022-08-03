package counter

import (
	"reflect"
	"testing"
)

type incorrectAVLStruct struct {
	key string
}

type correctAVLStruct struct {
	key string
}

func (c *correctAVLStruct) GetKey() string {
	return c.key
}

func TestAVLRequestCounterIncrementCounter(t *testing.T) {
	failedTests := []struct {
		name         string
		nodeToInsert *incorrectAVLStruct
	}{
		{
			name:         "the struct does not implemented binarySearchTree.AVLNodeData interface",
			nodeToInsert: &incorrectAVLStruct{key: "failed"},
		},
	}
	for _, test := range failedTests {
		t.Run(test.name, func(t *testing.T) {
			tree := NewAVLRequestCounter()
			if err := tree.IncrementCounter(test.nodeToInsert); err == nil {
				t.Errorf("should have an error")
			}
		})
	}
	successTests := []struct {
		name         string
		nodeToInsert []*correctAVLStruct
	}{
		{
			name: "insert multiple nodes",
			nodeToInsert: []*correctAVLStruct{
				{key: "abc"},
				{key: "def"},
				{key: "g"},
			},
		},
	}
	for _, test := range successTests {
		t.Run(test.name, func(t *testing.T) {
			tree := NewAVLRequestCounter()
			size := len(test.nodeToInsert)
			for x := 0; x < size; x++ {
				if err := tree.IncrementCounter(test.nodeToInsert[x]); err != nil {
					t.Errorf("should not have an error")
				}
			}
		})
	}
}

func TestAVLRequestCounterFindMostCalledRequests(t *testing.T) {
	tests := []struct {
		name            string
		nodeToIncrement []*correctAVLStruct
		want            []string
	}{
		{
			name:            "no request increment",
			nodeToIncrement: []*correctAVLStruct{},
			want:            []string{},
		},
		{
			name:            "one request increment",
			nodeToIncrement: []*correctAVLStruct{{key: "abc"}},
			want:            []string{"abc"},
		},
		{
			name: "multi requests increment, only one request is the most called",
			nodeToIncrement: []*correctAVLStruct{
				{key: "abc"}, {key: "def"}, {key: "abc"}, {key: "def"}, {key: "abc"}, {key: "lol"},
			},
			want: []string{"abc"},
		},
		{
			name: "multi requests increment, two requests is the most called",
			nodeToIncrement: []*correctAVLStruct{
				{key: "abc"}, {key: "def"}, {key: "abc"}, {key: "def"}, {key: "abc"}, {key: "lol"}, {key: "def"},
			},
			want: []string{"abc", "def"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := NewAVLRequestCounter()
			size := len(test.nodeToIncrement)
			for x := 0; x < size; x++ {
				if err := tree.IncrementCounter(test.nodeToIncrement[x]); err != nil {
					t.Errorf("should not have an error")
				}
			}
			mcr, err := tree.FindMostCalledRequests()
			if err != nil {
				t.Errorf("should not have an error")
			}
			size = len(mcr.([]*AVLNodeCounter))
			got := make([]string, size)
			for x := 0; x < size; x++ {
				got[x] = mcr.([]*AVLNodeCounter)[x].AVLNodeData.GetKey()
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("do not get the most called requests\nwant: %v\ngot: %v\n", test.want, got)
			}
		})
	}
}
