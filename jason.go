package jason

import (
	"encoding/json"
	"github.com/pkg/errors"
)

func NewJason(contents []byte) (Node, error) {
	if !json.Valid(contents) {
		return nil, errors.New("not json")
	}
	mapNode, err := newMapNodeContents(contents)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return mapNode, nil
}
