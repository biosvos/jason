package node

import "encoding/json"

var _ Node = &LazyNode{}

func NewLazyNode(parent *MapNode, key string, a any) *LazyNode {
	return &LazyNode{
		parent: parent,
		key:    key,
		data:   a,
	}
}

type LazyNode struct {
	parent *MapNode
	key    string
	data   any
}

func (l *LazyNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.data)
}

func (l *LazyNode) Number() (float64, error) {
	node := MakeFactory(l.data)
	l.parent.UpdateNode(l.key, node)
	return node.Number()
}

func (l *LazyNode) Bool() (bool, error) {
	node := MakeFactory(l.data)
	l.parent.UpdateNode(l.key, node)
	return node.Bool()
}

func (l *LazyNode) List() []Node {
	node := MakeFactory(l.data)
	l.parent.UpdateNode(l.key, node)
	return node.List()
}

func (l *LazyNode) Path(path string) []Node {
	node := MakeFactory(l.data)
	l.parent.UpdateNode(l.key, node)
	return node.Path(path)
}

func (l *LazyNode) Get(key string) (Node, error) {
	node := MakeFactory(l.data)
	l.parent.UpdateNode(l.key, node)
	return node.Get(key)
}

func (l *LazyNode) String() string {
	node := MakeFactory(l.data)
	l.parent.UpdateNode(l.key, node)
	return node.String()
}
