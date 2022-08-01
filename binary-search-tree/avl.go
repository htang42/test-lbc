package binarySearchTree

import (
	"sync"

	"github.com/htang42/test-lbc/utils"
)

type AVLNodeData interface {
	GetKey() string
}

type AVLNode struct {
	key    string
	left   *AVLNode
	right  *AVLNode
	height int
	Data   AVLNodeData
}

func (n *AVLNode) getHeight() int {
	if n == nil {
		return -1
	}
	return n.height
}

func (n *AVLNode) recalculateHeight() {
	n.height = 1 + utils.Max(n.left.getHeight(), n.right.getHeight())
}

func (n *AVLNode) calculateBalance() int {
	return n.right.getHeight() - n.left.getHeight()
}

func (n *AVLNode) rotateLeft() *AVLNode {
	rightChild := n.right

	n.right = rightChild.left
	rightChild.left = n

	n.recalculateHeight()
	rightChild.recalculateHeight()

	return rightChild
}

func (n *AVLNode) rotateRight() *AVLNode {
	leftChild := n.left

	n.left = leftChild.right
	leftChild.right = n

	n.recalculateHeight()
	leftChild.recalculateHeight()

	return leftChild
}

func (n *AVLNode) insert(nd AVLNodeData) *AVLNode {
	if n == nil {
		return &AVLNode{
			key:    nd.GetKey(),
			Data:   nd,
			left:   nil,
			right:  nil,
			height: 0,
		}
	}
	if nd.GetKey() < n.key {
		n.left = n.left.insert(nd)
	} else if nd.GetKey() > n.key {
		n.right = n.right.insert(nd)
	}
	n.recalculateHeight()
	return n.rebalance()
}

func (n *AVLNode) rebalance() *AVLNode {
	balance := n.calculateBalance()

	if balance < -1 {
		if n.left.calculateBalance() <= 0 {
			n = n.rotateRight()
		} else {
			n.left = n.left.rotateLeft()
			n = n.rotateRight()
		}
	}

	if balance > 1 {
		if n.right.calculateBalance() >= 0 {
			n = n.rotateLeft()
		} else {
			n.right = n.right.rotateRight()
			n = n.rotateLeft()
		}
	}
	return n
}

func (n *AVLNode) find(key string) *AVLNode {
	if n == nil {
		return nil
	}
	if key < n.key {
		return n.left.find(key)
	} else if key > n.key {
		return n.right.find(key)
	}
	return n
}

func (n *AVLNode) traverse(f func(*AVLNode)) {
	if n == nil {
		return
	}
	n.left.traverse(f)
	f(n)
	n.right.traverse(f)
}

type AVLTree struct {
	lock sync.RWMutex
	root *AVLNode
}

func (t *AVLTree) Insert(nd AVLNodeData) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.root = t.root.insert(nd)
}

func (t *AVLTree) Find(key string) *AVLNode {
	return t.root.find(key)
}

func (t *AVLTree) Traverse(f func(*AVLNode)) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.root.traverse(f)
}
