package jsons

import (
	"encoding/json"
	"fmt"

)

type JsonType interface {
	float64 | string | bool | map[string]any | []any
}

// Dig a value from a josn byte array by given cursors
func Cursor[A JsonType](data []byte, cursors ...any) (res A, err error) {
	var tmp any
	json.Unmarshal(data, &tmp)
	length := len(cursors)
	for i,j := range cursors {
			switch p := j.(type) {
			case string:
				switch v := tmp.(type) {
				case map[string]any:
					tmp = v[p]
				default:
					err = fmt.Errorf("expect map[string]any got %T at path %v", tmp, cursors[:i])
					return
				}
				if i == length-1 {
					var ok bool
					res, ok = tmp.(A); if !ok {
						err = fmt.Errorf("incorrect type %T at path %v\n", tmp, cursors[:i])
						return
					}
				}
			case int:
				switch v := (tmp).(type) {
				case []any:
					if(p < len(v)) {
						tmp = v[p]
					} else {
						err = fmt.Errorf("index out of range [%v/%v] at path %v", p, len(v), cursors[:i])
						return
					}
					
				default:
					err = fmt.Errorf("expect []any got %T at path %v", tmp, cursors[:i])
					return
				}
				if i == length-1 {
					var ok bool
					res, ok = tmp.(A); if !ok {
						err = fmt.Errorf("incorrect type %T at path %v\n", tmp, cursors[:i])
						return
					}
				}
			default:
				err = fmt.Errorf("cursor type can only be string or int, got %v at %d", j, i)
				return
			}
		
	}
	return
}

