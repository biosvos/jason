package node

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBoolNode_MarshalJSON(t *testing.T) {
	node := NewBoolNode(false)
	ret := node.String()
	require.Equal(t, "false", ret)
	marshalJSON, _ := node.MarshalJSON()
	require.EqualValues(t, "false", marshalJSON)
}
