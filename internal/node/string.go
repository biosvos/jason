package node

import "github.com/pkg/errors"

var _ Node = &StringNode{}

func NewStringNode(element string) *StringNode {
	return &StringNode{element: element}
}

type StringNode struct {
	element string
}

func (s *StringNode) List() []Node {
	return nil
}

func (s *StringNode) Path(path string) []Node {
	_ = path
	return nil
}

func (s *StringNode) Get(key string) (Node, error) {
	_ = key
	return nil, errors.New("failed to get node")
}

func (s *StringNode) String() string {
	return s.element
}
