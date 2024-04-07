package encoder

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/leehai1107/The-journey/internal/pkg/errors"
)

// EncryptionType represents the type of encryption algorithm used
type EncryptionType int

const (
	// AES is the type of encryption used
	AES EncryptionType = iota
	// RSA is another type of encryption
	RSA
)

// Key lengths for different encryption algorithms
const (
	bytesAES128 = 16
	bytesAES192 = 24
	byteAES256  = 32
)

// Encrypt encrypts the given data using the provided key and returns the encoded string.
func Encrypt(data []byte, key interface{}, encryptionType EncryptionType) (string, error) {
	switch encryptionType {
	case AES:
		// AES encryption
		return encryptAES(data, key.(string))
	case RSA:
		// RSA encryption
		return encryptRSA(data, key)
	default:
		return "", errors.MethodError.New()
	}
}

// Decrypt decrypts the given string using the provided key and returns the original data.
func Decrypt(encodedData string, key interface{}, encryptionType EncryptionType) ([]byte, error) {
	switch encryptionType {
	case AES:
		// AES decryption
		return decryptAES(encodedData, key.(string))
	case RSA:
		// RSA decryption
		return decryptRSA(encodedData, key)
	default:
		return nil, errors.MethodError.New()
	}
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
		return nil, errors.InvalidData.New()
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
		return "", errors.EncryptError.New()
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
		return nil, errors.InvalidData.New()
	}

	// Decode base64 encoded ciphertext
	ciphertext, err := base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		return nil, err
	}

	// Decrypt ciphertext using RSA private key
	return rsa.DecryptPKCS1v15(rand.Reader, rsaPrivKey, ciphertext)
}

// EncodeJSON encodes the given data in JSON format, then encrypts it using the provided key,
// and returns the encoded string.
func EncodeJSON(data interface{}, key interface{}, encryptionType EncryptionType) (string, error) {
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
			return "", errors.InvalidData.New()
		}
		encryptedData, err = Encrypt(jsonData, aesKey, AES)
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
		return "", errors.EncryptError.New()
	}

	return encryptedData, nil
}

// DecodeJSON decodes the given string using the provided key, then decrypts it from JSON format,
// and returns the original data.
func DecodeJSON(encodedData string, key interface{}, v interface{}, encryptionType EncryptionType) error {
	// Decrypt the encoded data
	var decryptedData []byte
	var err error
	switch encryptionType {
	case AES:
		// For AES decryption, use the provided key as a string
		aesKey, ok := key.(string)
		if !ok {
			return errors.InvalidData.New()
		}
		decryptedData, err = Decrypt(encodedData, aesKey, AES)
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
		return errors.DecryptError.New()
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
	key := make([]byte, keyLength)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}

	res := base64.URLEncoding.EncodeToString(key)
	if ValidateAESKey(res) {
		return res, nil
	}
	return "", errors.InvalidData.New()
}

// ValidateAESKey checks if the given key is a valid AES key
func ValidateAESKey(key string) bool {
	k := len([]byte(key))
	switch k {
	default:
		return false
	case bytesAES128, bytesAES192, byteAES256:
		return true
	}
}
