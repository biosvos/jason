package jason

import (
	"encoding/json"
	"github.com/pkg/errors"
)

func newMapNode(elements map[string]any) *MapNode {
	ret := MapNode{
		elements: map[string]Node{},
	}
	for key, value := range elements {
		ret.AddNode(key, value)
	}
	return &ret
}

func newMapNodeContents(contents []byte) (*MapNode, error) {
	var elements map[string]any
	err := json.Unmarshal(contents, &elements)
	if err != nil {
		// 타입이 다를 수 있다.
		return nil, errors.WithStack(err)
	}
	return newMapNode(elements), nil
}

var _ Node = &MapNode{}

type MapNode struct {
	elements map[string]Node
}

func (m *MapNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&m.elements)
}

func (m *MapNode) Number() (float64, error) {
	return 0, errors.New("failed to get number")
}

func (m *MapNode) Bool() (bool, error) {
	return false, errors.New("failed to get bool")
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
	m.elements[key] = newLazyNode(m, key, value)
}

func (m *MapNode) String() string {
	marshal, _ := json.Marshal(m.elements) //nolint:errchkjson
	return string(marshal)
}

func (m *MapNode) Path(path string) []Node {
	return nodePath(m, path)
}

func (m *MapNode) Get(key string) (Node, error) {
	ret, ok := m.elements[key]
	if !ok {
		return nil, errors.New("not found")
	}
	return ret, nil
}
