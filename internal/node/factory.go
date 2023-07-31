package node

func MakeFactory(a any) Node {
	switch v := a.(type) {
	case string:
		return NewStringNode(v)
	case []interface{}:
		return NewSliceNode(v)
	default:
		panic(v)
	}
}
