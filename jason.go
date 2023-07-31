package jason

import (
	"encoding/json"
	"github.com/biosvos/jason/internal/node"
	"github.com/pkg/errors"
)

func NewJason(contents []byte) (node.Node, error) {
	if !json.Valid(contents) {
		return nil, errors.New("not json")
	}
	mapNode, err := node.NewMapNode(contents)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return mapNode, nil
}
