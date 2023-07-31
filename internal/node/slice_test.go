package node

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSliceNode_MarshalJSON(t *testing.T) {
	node := NewSliceNode([]any{1, 2, 3})
	require.Equal(t, "[1,2,3]", node.String())
	json, _ := node.MarshalJSON()
	require.EqualValues(t, "[1,2,3]", string(json))
}
