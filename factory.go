package jason

func makeFactory(a any) Node {
	switch v := a.(type) {
	case string:
		return newStringNode(v)
	case []any:
		return newSliceNode(v)
	case bool:
		return newBoolNode(v)
	case int, int64, int32, int16, uint, uint16, uint32, uint64, float64, float32:
		return numbericFactory(v)
	case map[string]any:
		return newMapNode(v)
	default:
		panic(v)
	}
}

func numbericFactory(a any) *NumberNode {
	switch v := a.(type) {
	case int:
		return newNumberNode(float64(v))
	case int64:
		return newNumberNode(float64(v))
	case int32:
		return newNumberNode(float64(v))
	case int16:
		return newNumberNode(float64(v))
	case uint:
		return newNumberNode(float64(v))
	case uint16:
		return newNumberNode(float64(v))
	case uint32:
		return newNumberNode(float64(v))
	case uint64:
		return newNumberNode(float64(v))
	case float64:
		return newNumberNode(float64(v))
	case float32:
		return newNumberNode(float64(v))
	default:
		panic(v)
	}
}
