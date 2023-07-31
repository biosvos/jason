package jason

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumberNode_MarshalJSON(t *testing.T) {
	node := newNumberNode(1)
	require.Equal(t, "1", node.String())
	json, _ := node.MarshalJSON()
	require.EqualValues(t, "1", json)
}
