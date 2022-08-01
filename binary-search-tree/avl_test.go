package binarySearchTree

import (
	"reflect"
	"testing"
)

type testAVLNodeData struct {
	key string
}

func (n testAVLNodeData) GetKey() string {
	return n.key
}

func compareTree(n1, n2 *AVLNode) bool {
	if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
		return false
	}
	if n1 == nil && n2 == nil {
		return true
	}
	if n1.key != n2.key {
		return false
	}
	if !compareTree(n1.left, n2.left) {
		return false
	}
	if !compareTree(n1.right, n2.right) {
		return false
	}
	return true
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name         string
		nodeToInsert []*testAVLNodeData
		wantedTree   *AVLTree
	}{
		{
			name: "first test multiple inserts",
			nodeToInsert: []*testAVLNodeData{
				{key: "1"}, {key: "2"}, {key: "3"}, {key: "4"}, {key: "5"}, {key: "6"}, {key: "7"},
			},
			wantedTree: &AVLTree{
				root: &AVLNode{
					key: "4",
					left: &AVLNode{
						key:   "2",
						left:  &AVLNode{key: "1"},
						right: &AVLNode{key: "3"},
					},
					right: &AVLNode{
						key:   "6",
						left:  &AVLNode{key: "5"},
						right: &AVLNode{key: "7"},
					},
				},
			},
		},
		{
			name: "second test multiple inserts",
			nodeToInsert: []*testAVLNodeData{
				{key: "5"}, {key: "6"}, {key: "4"}, {key: "3"}, {key: "1"}, {key: "7"}, {key: "2"},
			},
			wantedTree: &AVLTree{
				root: &AVLNode{
					key: "5",
					left: &AVLNode{
						key: "3",
						left: &AVLNode{
							key:   "1",
							right: &AVLNode{key: "2"},
						},
						right: &AVLNode{key: "4"},
					},
					right: &AVLNode{
						key:   "6",
						right: &AVLNode{key: "7"},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := &AVLTree{}
			size := len(test.nodeToInsert)
			for x := 0; x < size; x++ {
				tree.Insert(test.nodeToInsert[x])
			}
			if !compareTree(tree.root, test.wantedTree.root) {
				t.Errorf("not same tree")
			}
		})
	}

}

func TestTraverse(t *testing.T) {
	tests := []struct {
		name         string
		nodeToInsert []*testAVLNodeData
		wantedOrder  []string
	}{
		{
			name: "insert nodes in order",
			nodeToInsert: []*testAVLNodeData{
				{key: "1"}, {key: "2"}, {key: "3"}, {key: "4"}, {key: "5"}, {key: "6"}, {key: "7"},
			},
			wantedOrder: []string{"1", "2", "3", "4", "5", "6", "7"},
		},
		{
			name: "insert nodes in disorder",
			nodeToInsert: []*testAVLNodeData{
				{key: "6"}, {key: "4"}, {key: "1"}, {key: "5"}, {key: "2"}, {key: "7"}, {key: "3"},
			},
			wantedOrder: []string{"1", "2", "3", "4", "5", "6", "7"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := &AVLTree{}
			size := len(test.nodeToInsert)
			for x := 0; x < size; x++ {
				tree.Insert(test.nodeToInsert[x])
			}
			gotOrder := make([]string, 0)
			tree.Traverse(func(n *AVLNode) {
				if n != nil {
					gotOrder = append(gotOrder, n.key)
				}
			})
			if !reflect.DeepEqual(test.wantedOrder, gotOrder) {
				t.Errorf("not the same order")
			}
		})
	}
}
