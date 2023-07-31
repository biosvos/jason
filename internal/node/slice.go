package node

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func NewSliceNode(slice []any) *SliceNode {
	var elements []Node
	for _, item := range slice {
		elements = append(elements, MakeFactory(item))
	}
	return &SliceNode{
		elements: elements,
	}
}

var _ Node = &SliceNode{}

type SliceNode struct {
	elements []Node
}

func (s *SliceNode) Number() (float64, error) {
	return 0, errors.New("failed to get number")
}

func (s *SliceNode) Bool() (bool, error) {
	return false, errors.New("failed to get bool")
}

func (s *SliceNode) List() []Node {
	return s.elements
}

func (s *SliceNode) Path(path string) []Node {
	return Path(s, path)
}

func (s *SliceNode) Get(key string) (Node, error) {
	index64, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	index := int(index64)
	if index < 0 || len(s.elements) <= index {
		return nil, errors.New("index overflow")
	}
	return s.elements[index], nil
}

func (s *SliceNode) String() string {
	var slice []string
	for _, element := range s.elements {
		slice = append(slice, element.String())
	}
	return strings.Join(slice, ",")
}
