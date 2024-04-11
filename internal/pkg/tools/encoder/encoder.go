package encoder

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"strings"
)

// EncryptionType represents the type of encryption algorithm used
type EncryptionType int

const (
	// AES is the type of encryption used
	AES EncryptionType = iota
	// RSA is another type of encryption
	RSA
	// DES is another type of encryption
	DES
)

// Key lengths for different encryption algorithms
const (
	bytesDES    = 8
	bytesAES128 = 16
	bytesAES192 = 24
	bytesAES256 = 32
)

// Base64StdEncode encode string with base64 encoding.
// Play: https://go.dev/play/p/VOaUyQUreoK
func Base64StdEncode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64StdDecode decode a base64 encoded string.
// Play: https://go.dev/play/p/RWQylnJVgIe
func Base64StdDecode(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}

// BasicEncodeToJSON encodes data to JSON string
func BasicEncodeToJSON(data interface{}) (string, error) {
	// Encode data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// BasicDecodeFromJSON decodes JSON string to data
func BasicDecodeFromJSON(jsonString string) (map[string]interface{}, error) {
	// Decode JSON string to data
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Encrypt encrypts the given data using the provided key and encryption algorithm.
func Encrypt(data []byte, key interface{}, encryptionType EncryptionType) (string, error) {
	switch encryptionType {
	case AES:
		// AES encryption
		return encryptAES(data, key.(string))
	case DES:
		// DES encryption
		return encryptDES(data, key.(string))
	case RSA:
		// RSA encryption
		return encryptRSA(data, key)
	default:
		return "", errors.New("invalid encryption type")
	}
}

// Decrypt decrypts the given string using the provided key and encryption algorithm.
func Decrypt(encodedData string, key interface{}, encryptionType EncryptionType) ([]byte, error) {
	switch encryptionType {
	case AES:
		// AES decryption
		return decryptAES(encodedData, key.(string))
	case DES:
		// DES decryption
		return decryptDES(encodedData, key.(string))
	case RSA:
		// RSA decryption
		return decryptRSA(encodedData, key)
	default:
		return nil, errors.New("invalid encryption type")
	}
}

// EncodeJSONWithKey encodes the given data in JSON format, then encrypts it using the provided key,
// and returns the encoded string.
func EncodeJSONWithKey(data interface{}, key interface{}, encryptionType EncryptionType) (string, error) {
	// Marshal the data into JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Encrypt the JSON data
	var encryptedData string
	switch encryptionType {
	case AES:
		// For AES encryption, use the provided key as a string
		aesKey, ok := key.(string)
		if !ok {
			return "", errors.New("invalid key type")
		}
		encryptedData, err = Encrypt(jsonData, aesKey, AES)
		if err != nil {
			return "", err
		}
	case DES:
		// For DES encryption, use the provided key as a string
		desKey, ok := key.(string)
		if !ok {
			return "", errors.New("invalid key type")
		}
		encryptedData, err = encryptDES(jsonData, desKey)
		if err != nil {
			return "", err
		}
	case RSA:
		// For RSA encryption, use the provided key directly
		encryptedData, err = encryptRSA(jsonData, key)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("invalid encryption type")
	}

	return encryptedData, nil
}

// DecodeJSONWithKey decodes the given string using the provided key, then decrypts it from JSON format,
// and returns the original data.
func DecodeJSONWithKey(encodedData string, key interface{}, v interface{}, encryptionType EncryptionType) error {
	// Decrypt the encoded data
	var decryptedData []byte
	var err error
	switch encryptionType {
	case AES:
		// For AES decryption, use the provided key as a string
		aesKey, ok := key.(string)
		if !ok {
			return errors.New("invalid key type")
		}
		decryptedData, err = Decrypt(encodedData, aesKey, AES)
		if err != nil {
			return err
		}
	case DES:
		// For DES decryption, use the provided key as a string
		desKey, ok := key.(string)
		if !ok {
			return errors.New("invalid key type")
		}
		decryptedData, err = decryptDES(encodedData, desKey)
		if err != nil {
			return err
		}
	case RSA:
		// For RSA decryption, use the provided key directly
		decryptedData, err = decryptRSA(encodedData, key)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid encryption type")
	}

	// Unmarshal the decrypted data from JSON format
	if err := json.Unmarshal(decryptedData, &v); err != nil {
		return err
	}

	return nil
}

// GenerateRSAKeyPair generates a new RSA key pair with the specified key size.
func GenerateRSAKeyPair(keySize int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// GenerateAESKey generates a random AES key of the specified length
func GenerateAESKey(keyLength int) (string, error) {
	switch keyLength {
	default:
		return "", errors.New("invalid key length")
	case bytesAES128, bytesAES192, bytesAES256:
		key := make([]byte, keyLength)
		if _, err := rand.Read(key); err != nil {
			return "", err
		}

		return customEncode(key), nil
	}
}

// GenerateDESKey generates a random DES key of the specified length
func GenerateDESKey() (string, error) {
	key := make([]byte, bytesDES)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}

	return customEncode(key), nil
}

// ValidateAESKey checks if the given key is a valid AES key
func ValidateAESKey(key string) bool {
	k := len([]byte(key))
	switch k {
	default:
		return false
	case bytesAES128, bytesAES192, bytesAES256:
		return true
	}
}

// ValidateDESKey checks if the given key is a valid DES key
func ValidateDESKey(key string) bool {
	k := len([]byte(key))
	switch k {
	default:
		return false
	case bytesDES:
		return true
	}
}

// customEncode encodes a byte slice to a custom string without increasing its length
func customEncode(data []byte) string {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+-_=*./<>?(&^%$#@!~`)[]{}"
	var encoded strings.Builder
	for _, b := range data {
		encoded.WriteByte(alphabet[b%byte(len(alphabet))])
	}
	return encoded.String()
}

// Encrypt data using AES algorithm
func encryptAES(data []byte, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt data using AES algorithm
func decryptAES(encodedData string, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	ciphertext, err := base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext is too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

// Encrypt data using RSA algorithm
func encryptRSA(data []byte, key interface{}) (string, error) {
	// Convert the key to *rsa.PublicKey
	rsaPubKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return "", errors.New("invalid key type")
	}

	// Encrypt data using RSA public key
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPubKey, data)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt data using RSA algorithm
func decryptRSA(encodedData string, key interface{}) ([]byte, error) {
	// Convert the key to *rsa.PrivateKey
	rsaPrivKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("invalid key type")
	}

	// Decode base64 encoded ciphertext
	ciphertext, err := base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		return nil, err
	}

	// Decrypt ciphertext using RSA private key
	return rsa.DecryptPKCS1v15(rand.Reader, rsaPrivKey, ciphertext)
}

// Encrypt data using DES algorithm
func encryptDES(data []byte, key string) (string, error) {
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, des.BlockSize+len(data))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[des.BlockSize:], data)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt data using DES algorithm
func decryptDES(encodedData string, key string) ([]byte, error) {
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	ciphertext, err := base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext is too short")
	}
	iv := ciphertext[:des.BlockSize]
	ciphertext = ciphertext[des.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
