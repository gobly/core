package json

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Decimal struct {
	primitive.Decimal128
}

func (d Decimal) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Decimal) UnmarshalJSON(b []byte) error {
	var decimalNumber json.Number
	err := json.Unmarshal(b, &decimalNumber)
	if err == nil {
		d.Decimal128, err = primitive.ParseDecimal128(decimalNumber.String())
	}

	return err
}
