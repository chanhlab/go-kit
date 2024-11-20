package uuid

import "github.com/lithammer/shortuuid/v4"

// ShortUUID generates new UUID
func ShortUUID() string {
	return shortuuid.New()
}
