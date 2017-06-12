package json

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

type Decimal struct {
	bson.Decimal128
}

func (d Decimal) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Decimal) UnmarshalJSON(b []byte) error {
	var decimalNumber json.Number
	err := json.Unmarshal(b, &decimalNumber)
	if err == nil {
		d.Decimal128, err = bson.ParseDecimal128(decimalNumber.String())
	}

	return err
}
