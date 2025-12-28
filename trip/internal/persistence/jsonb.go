package persistence

import (
	"database/sql/driver"
	"encoding/json"
)

type JSONB map[string]any

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(value any) error {
	return json.Unmarshal(value.([]byte), j)
}
