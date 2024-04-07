package encoder

import (
	"encoding/base64"
	"encoding/json"

	"github.com/leehai1107/The-journey/internal/pkg/errors"
)

// Encode encodes the given data using the provided key and returns the encoded string.
func Encode(data []byte, key string) (string, error) {
	// Append the key to the data before encoding
	data = append(data, []byte(key)...)

	// Encode the data using base64
	encodedData := base64.StdEncoding.EncodeToString(data)

	return encodedData, nil
}

// Decode decodes the given string using the provided key and returns the original data.
func Decode(encodedData string, key string) ([]byte, error) {
	// Decode the base64 encoded string
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return nil, err
	}

	// Check if the decoded data has enough length to contain the key
	if len(decodedData) <= len(key) {
		return nil, errors.InvalidData.New()
	}

	// Extract the original data by removing the key
	originalData := decodedData[:len(decodedData)-len(key)]

	return originalData, nil
}

// EncodeJSON encodes the given data in JSON format, then encodes it using the provided key,
// and returns the encoded string.
func EncodeJSON(data interface{}, key string) (string, error) {
	// Marshal the data into JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Encode the JSON data using base64
	encodedData, err := Encode(jsonData, key)
	if err != nil {
		return "", err
	}

	return encodedData, nil
}

// DecodeJSON decodes the given string using the provided key, then decodes it from JSON format,
// and returns the original data.
func DecodeJSON(encodedData string, key string, v interface{}) error {
	// Decode the encoded data using base64
	decodedData, err := Decode(encodedData, key)
	if err != nil {
		return err
	}

	// Unmarshal the decoded data from JSON format
	if err := json.Unmarshal(decodedData, &v); err != nil {
		return err
	}

	return nil
}
