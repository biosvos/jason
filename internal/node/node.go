package node

type Node interface {
	Path(path string) []Node
	Get(key string) (Node, error)
	List() []Node
	String() string
	Number() (float64, error)
	Bool() (bool, error)
}
