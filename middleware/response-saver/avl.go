package responseSaver

import (
	"fmt"

	binarySearchTree "github.com/htang42/test-lbc/binary-search-tree"
)

type AVLNodeResponseSaver struct {
	binarySearchTree.AVLNodeData
	Response interface{}
}

type AVLResponseSaver struct {
	*binarySearchTree.AVLTree
}

func (rs *AVLResponseSaver) SetResponse(request interface{}, response interface{}) error {
	n, ok := request.(binarySearchTree.AVLNodeData)
	if !ok {
		return fmt.Errorf("the request struct does not implement binarySearchTree.AVLNodeData interface")
	}
	node := rs.Find(n.GetKey())
	if node == nil {
		rs.Insert(&AVLNodeResponseSaver{AVLNodeData: n, Response: response})
	}
	return nil
}

func (rs *AVLResponseSaver) GetResponse(request interface{}) (interface{}, error) {
	n, ok := request.(binarySearchTree.AVLNodeData)
	if !ok {
		return nil, fmt.Errorf("the request struct does not implement binarySearchTree.AVLNodeData interface")
	}
	node := rs.Find(n.GetKey())
	if node == nil {
		return nil, nil
	}
	return node.Data.(*AVLNodeResponseSaver).Response, nil
}

func NewAVLResponseSaver() *AVLResponseSaver {
	return &AVLResponseSaver{
		&binarySearchTree.AVLTree{},
	}
}
