package jason

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewJason(t *testing.T) {
	t.Run("failed", func(t *testing.T) {
		root, err := NewJason([]byte("{invalid}"))
		require.Error(t, err)
		require.Nil(t, root)
	})
	t.Run("", func(t *testing.T) {
		root, err := NewJason([]byte("{}"))
		require.NoError(t, err)
		require.NotNil(t, root)
	})
}

func TestGet(t *testing.T) {
	t.Run("failed", func(t *testing.T) {
		root, _ := NewJason([]byte("{}"))

		get, err := root.Get("A")

		require.Error(t, err)
		require.Nil(t, get)
	})
	t.Run("", func(t *testing.T) {
		root, _ := NewJason([]byte(`{"A":"HI"}`))

		get, err := root.Get("A")

		require.NoError(t, err)
		require.NotNil(t, get)
	})
}

func TestJson(t *testing.T) {
	t.Run("", func(t *testing.T) {
		root, _ := NewJason([]byte(`{"A":"HI"}`))

		t.Log(root.String())
	})
}

func TestString(t *testing.T) {
	root, _ := NewJason([]byte(`{"A":"HI"}`))
	get, _ := root.Get("A")

	ret := get.String()

	require.Equal(t, "HI", ret)
}

func TestPath(t *testing.T) {
	root, _ := NewJason([]byte(`{"A":["one", "two"]}`))

	nodes := root.Path("A.0")

	require.Len(t, nodes, 1)
	require.Equal(t, "one", nodes[0].String())
}
