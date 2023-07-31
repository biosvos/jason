package node

func MakeFactory(a any) Node {
	switch v := a.(type) {
	case string:
		return NewStringNode(v)
	case []any:
		return NewSliceNode(v)
	case bool:
		return NewBoolNode(v)
	case int, int64, int32, int16, uint, uint16, uint32, uint64, float64, float32:
		return numbericFactory(v)
	default:
		panic(v)
	}
}

func numbericFactory(a any) *NumberNode {
	switch v := a.(type) {
	case int:
		return NewNumberNode(float64(v))
	case int64:
		return NewNumberNode(float64(v))
	case int32:
		return NewNumberNode(float64(v))
	case int16:
		return NewNumberNode(float64(v))
	case uint:
		return NewNumberNode(float64(v))
	case uint16:
		return NewNumberNode(float64(v))
	case uint32:
		return NewNumberNode(float64(v))
	case uint64:
		return NewNumberNode(float64(v))
	case float64:
		return NewNumberNode(float64(v))
	case float32:
		return NewNumberNode(float64(v))
	default:
		panic(v)
	}
}
