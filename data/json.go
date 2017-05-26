package json

import (
	"encoding/json"
	"io"
)

func Encode(out io.Writer, model interface{}) error {
	return json.NewEncoder(out).Encode(model)
}

func Decode(in io.Reader, model interface{}) error {
	return json.NewDecoder(in).Decode(model)
}

func Marshal(model interface{}) ([]byte, error) {
	return json.Marshal(model)
}
