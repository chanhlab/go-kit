package uuid

import "github.com/lithammer/shortuuid/v3"

// NewUUID generates new UUID
func NewUUID() string {
	return shortuuid.New()
}
