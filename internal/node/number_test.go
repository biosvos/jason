package node

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumberNode_MarshalJSON(t *testing.T) {
	node := NewNumberNode(1)
	require.Equal(t, "1", node.String())
	json, _ := node.MarshalJSON()
	require.EqualValues(t, "1", json)
}
