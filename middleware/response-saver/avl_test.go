package responseSaver

import (
	"reflect"
	"testing"
)

type incorrectAVLStruct struct {
	key string
}

type correctAVLStruct struct {
	key    string
	result string
}

func (c *correctAVLStruct) GetKey() string {
	return c.key
}

func TestAVLResponseSaverSetResponse(t *testing.T) {
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
			tree := NewAVLResponseSaver()
			if err := tree.SetResponse(test.nodeToInsert, "failed"); err == nil {
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
			tree := NewAVLResponseSaver()
			size := len(test.nodeToInsert)
			for x := 0; x < size; x++ {
				if err := tree.SetResponse(test.nodeToInsert[x], "success"); err != nil {
					t.Errorf("should not have an error")
				}
			}
		})
	}
}

func TestAVLResponseSaverGetResponse(t *testing.T) {
	tests := []struct {
		name       string
		nodeToSave []*correctAVLStruct
	}{
		{
			name: "one node insert",
			nodeToSave: []*correctAVLStruct{
				{key: "abc", result: "result_abc"},
			},
		},
		{
			name: "many nodes insert",
			nodeToSave: []*correctAVLStruct{
				{key: "abc", result: "123"},
				{key: "def", result: "456"},
				{key: "ghi", result: "789"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := NewAVLResponseSaver()
			size := len(test.nodeToSave)
			for x := 0; x < size; x++ {
				if err := tree.SetResponse(test.nodeToSave[x], test.nodeToSave[x].result); err != nil {
					t.Errorf("should not have an error")
				}
			}
			for x := 0; x < size; x++ {
				result, err := tree.GetResponse(test.nodeToSave[x])
				if err != nil {
					t.Errorf("should not have an error")
				}
				if !reflect.DeepEqual(result, test.nodeToSave[x].result) {
					t.Errorf("do not get the right response")
				}
			}
		})
	}
}
