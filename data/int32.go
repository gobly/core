package json

import (
	"encoding/json"
	"strconv"
)

type Int32 int32

func (o *Int32) UnmarshalJSON(b []byte) error {
	var number json.Number
	var i int64

	err := json.Unmarshal(b, &number)
	if err == nil {
		i, err = strconv.ParseInt(string(number), 10, 32)
		*o = Int32(i)
	}

	return err
}
