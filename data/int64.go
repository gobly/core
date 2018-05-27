package json

import (
	"encoding/json"
)

type Int64 int64

func (o *Int64) UnmarshalJSON(b []byte) error {
	var number json.Number
	var i int64

	err := json.Unmarshal(b, &number)
	if err == nil {
		i, err = number.Int64()
		if err == nil {
			*o = Int64(i)
		}
	}

	return err
}
