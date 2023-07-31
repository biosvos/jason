package node

import (
	"encoding/json"
	"github.com/pkg/errors"
)

func NewMapNode(contents []byte) (*MapNode, error) {
	var elements map[string]any
	err := json.Unmarshal(contents, &elements)
	if err != nil {
		// 타입이 다를 수 있다.
		return nil, errors.WithStack(err)
	}
	ret := MapNode{
		elements: map[string]Node{},
	}
	for key, value := range elements {
		ret.AddNode(key, value)
	}
	return &ret, nil
}

var _ Node = &MapNode{}

type MapNode struct {
	elements map[string]Node
}

func (m *MapNode) List() []Node {
	var ret []Node
	for _, node := range m.elements {
		ret = append(ret, node)
	}
	return ret
}

func (m *MapNode) UpdateNode(key string, value Node) {
	m.elements[key] = value
}

func (m *MapNode) AddNode(key string, value any) {
	m.elements[key] = NewLazyNode(m, key, value)
}

func (m *MapNode) String() string {
	marshal, _ := json.Marshal(m.elements) //nolint:errchkjson
	return string(marshal)
}

func (m *MapNode) Path(path string) []Node {
	return Path(m, path)
}

func (m *MapNode) Get(key string) (Node, error) {
	ret, ok := m.elements[key]
	if !ok {
		return nil, errors.New("not found")
	}
	return ret, nil
}
