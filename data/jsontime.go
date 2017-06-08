package json

import (
	"encoding/json"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Format(time.RFC3339))
}
