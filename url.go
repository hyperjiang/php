package php

import (
	"encoding/base64"
)

// Base64Encode encodes data with MIME base64
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Base64Decode decodes data encoded with MIME base64.
// Returns error if data is in invalid format.
func Base64Decode(data string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(data)

	return string(b), err
}
