package cryptoUtils

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
)

func GenerateBytes(max uint) []byte {
	var slice = make([]byte, max)
	_, err := rand.Read(slice)
	if err != nil {
		return nil
	}
	return slice
}

func GeneratePrivate(bits uint) *rsa.PrivateKey {
	priv, err := rsa.GenerateKey(rand.Reader, int(bits))
	if err != nil {
		return nil
	}
	return priv
}

func HashPublic(pub *rsa.PublicKey) string {
	return Base64Encode(HashSum([]byte(StringPublic(pub))))
}

func HashSum(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func ParsePrivate(privData string) *rsa.PrivateKey {
	priv, err := x509.ParsePKCS1PrivateKey(Base64Decode(privData))
	if err != nil {
		return nil
	}
	return priv
}

func ParsePublic(publicData string) *rsa.PublicKey {
	pub, err := x509.ParsePKCS1PublicKey(Base64Decode(publicData))
	if err != nil {
		return nil
	}
	return pub
}

func StringPrivate(priv *rsa.PrivateKey) string {
	return Base64Encode(x509.MarshalPKCS1PrivateKey(priv))
}

func StringPublic(pub *rsa.PublicKey) string {
	return Base64Encode(x509.MarshalPKCS1PublicKey(pub))
}

func EncryptRSA(pub *rsa.PublicKey, datai []byte) []byte {
	datao, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, datai, nil)
	if err != nil {
		return nil
	}
	return datao
}

func DecryptRSA(priv *rsa.PrivateKey, datai []byte) []byte {
	datao, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, datai, nil)
	if err != nil {
		return nil
	}
	return datao
}

func EncryptAES(key, data []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		println(err.Error())
		return nil
	}
	blockSize := block.BlockSize()
	data = paddingPKCS5(data, blockSize)
	cipherText := make([]byte, blockSize+len(data))
	iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], data)
	return cipherText
}

func DecryptAES(key, data []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	blockSize := block.BlockSize()
	if len(data) < blockSize {
		return nil
	}
	iv := data[:blockSize]
	data = data[blockSize:]
	if len(data)%blockSize != 0 {
		return nil
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)
	return unpaddingPKCS5(data)
}

func paddingPKCS5(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func unpaddingPKCS5(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return nil
	}
	unpadding := int(data[length-1])
	if length < unpadding {
		return nil
	}
	return data[:(length - unpadding)]
}

func Sign(priv *rsa.PrivateKey, data []byte) []byte {
	sign, err := rsa.SignPSS(rand.Reader, priv, crypto.SHA256, data, nil)
	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}
	return sign
}

func Verify(pub *rsa.PublicKey, data, sign []byte) error {
	return rsa.VerifyPSS(pub, crypto.SHA256, data, sign, nil)
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(data string) []byte {
	result, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil
	}
	return result
}

func PaddingPassword(password string) ([]byte, error) {
	psswd := []byte(password)
	pass_len := len(password)
	padding := 16 - pass_len
	if pass_len > 16 {
		return nil, errors.New("password len is more then 16 char")
	}
	padPass := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(psswd, padPass...), nil
}

func GetFileBytes(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

func GetHashFromFile(path string) (string, error) {
	data, err := GetFileBytes(path)
	if err != nil {
		return "", err
	}
	return Base64Encode(HashSum(data)), err
}
