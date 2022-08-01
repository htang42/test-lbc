package counter

import (
	"fmt"
	"sync"

	binarySearchTree "github.com/htang42/test-lbc/binary-search-tree"
)

type AVLNodeCounter struct {
	binarySearchTree.AVLNodeData
	Count int
	lock  sync.RWMutex
}

type AVLRequestCounter struct {
	*binarySearchTree.AVLTree
}

func (rc *AVLRequestCounter) IncrementCounter(i interface{}) error {
	n, ok := i.(binarySearchTree.AVLNodeData)
	if !ok {
		return fmt.Errorf("the request struct does not implement binarySearchTree.AVLNodeData interface")
	}
	node := rc.Find(n.GetKey())
	if node == nil {
		rc.Insert(&AVLNodeCounter{AVLNodeData: n, Count: 1})
		return nil
	}
	nc := node.Data.(*AVLNodeCounter)
	nc.lock.Lock()
	defer nc.lock.Unlock()
	nc.Count++
	return nil
}

func (rc *AVLRequestCounter) FindMostCalledRequests() (interface{}, error) {
	mostCalledRequest := make([]*AVLNodeCounter, 0)
	countRef := 0
	rc.Traverse(func(n *binarySearchTree.AVLNode) {
		nd := n.Data.(*AVLNodeCounter)
		if countRef < nd.Count {
			mostCalledRequest = []*AVLNodeCounter{nd}
			countRef = nd.Count
		} else if countRef == nd.Count {
			mostCalledRequest = append(mostCalledRequest, nd)
		}
	})
	return mostCalledRequest, nil
}

func NewAVLRequestCounter() *AVLRequestCounter {
	return &AVLRequestCounter{
		&binarySearchTree.AVLTree{},
	}
}
