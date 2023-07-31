package node

import (
	"github.com/pkg/errors"
	"strconv"
)

var _ Node = &NumberNode{}

func NewNumberNode(element float64) *NumberNode {
	return &NumberNode{
		element: element,
	}
}

type NumberNode struct {
	element float64
}

func (n *NumberNode) Path(path string) []Node {
	_ = path
	return nil
}

func (n *NumberNode) Get(key string) (Node, error) {
	_ = key
	return nil, errors.New("failed to get key")
}

func (n *NumberNode) List() []Node {
	return nil
}

func (n *NumberNode) String() string {
	return strconv.FormatFloat(n.element, 'f', -1, 64)
}

func (n *NumberNode) Number() (float64, error) {
	return n.element, nil
}

func (n *NumberNode) Bool() (bool, error) {
	return false, errors.New("failed to get bool")
}
