package common

import "encoding/base64"

// HasedKey hashes the key using base64 encoding
func HashKey(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// DecodeHashedKey decodes the hased key
func DecodeHashedKey(input string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
