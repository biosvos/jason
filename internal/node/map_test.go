package node

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMapNode_MarshalJSON(t *testing.T) {
	node := NewMapNode(map[string]any{
		"A": "C",
		"B": []string{"D", "E"},
	})
	require.Equal(t, `{"A":"C","B":["D","E"]}`, node.String())
	json, _ := node.MarshalJSON()
	require.EqualValues(t, `{"A":"C","B":["D","E"]}`, json)
}
