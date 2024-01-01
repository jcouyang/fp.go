package jsons

import (
	"encoding/json"

	"github.com/jcouyang/fp.go/maps"
)

type JsonType interface {
	float64 | string | bool | map[string]any | []any
}

// Dig a value from a josn byte array by given cursors
func Dig[A JsonType](data []byte, cursors ...any) (res A, err error) {
	var tmp any
	json.Unmarshal(data, &tmp)
	return maps.Dig[A](tmp, cursors...)
}
