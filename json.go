package php

import (
	"encoding/json"
)

// JSONEncode returns a string containing the JSON representation of the supplied value
func JSONEncode(value any) (string, error) {
	b, err := json.Marshal(value)
	return string(b), err
}

// JSONDecode decodes a JSON string and stores the result in the value pointed to by v
//
// Be aware that this is different with php json_decode
func JSONDecode(jsonStr string, v any) error {
	return json.Unmarshal([]byte(jsonStr), v)
}
