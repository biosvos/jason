package node

import (
	"encoding/json"
	"errors"
	"strconv"
)

var _ Node = &BoolNode{}

func NewBoolNode(element bool) *BoolNode {
	return &BoolNode{
		element: element,
	}
}

type BoolNode struct {
	element bool
}

func (b *BoolNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&b.element)
}

func (b *BoolNode) Path(path string) []Node {
	_ = path
	return nil
}

func (b *BoolNode) Get(key string) (Node, error) {
	_ = key
	return nil, errors.New("not found key")
}

func (b *BoolNode) List() []Node {
	return nil
}

func (b *BoolNode) String() string {
	return strconv.FormatBool(b.element)
}

func (b *BoolNode) Number() (float64, error) {
	if b.element {
		return 1, nil
	}
	return 0, nil
}

func (b *BoolNode) Bool() (bool, error) {
	return b.element, nil
}
