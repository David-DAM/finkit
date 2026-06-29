package cache

import (
	"encoding/json"
	"time"
)

type WrapperGet struct {
	Value     json.RawMessage `json:"value"`
	ExpiresAt time.Time       `json:"expires_at"`
}

type WrapperSet struct {
	Value     any       `json:"value"`
	ExpiresAt time.Time `json:"expires_at"`
}
