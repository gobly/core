package data

import (
	"encoding/json"
	"io"
)

func Encode(out io.Writer, data interface{}) {
	json.NewEncoder(out).Encode(data)
}
