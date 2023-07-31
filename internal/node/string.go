package node

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
)

var _ Node = &StringNode{}

func NewStringNode(element string) *StringNode {
	return &StringNode{element: element}
}

type StringNode struct {
	element string
}

func (s *StringNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&s.element)
}

func (s *StringNode) Number() (float64, error) {
	float, err := strconv.ParseFloat(s.element, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return float, nil
}

func (s *StringNode) Bool() (bool, error) {
	ret, err := strconv.ParseBool(s.element)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return ret, nil
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
