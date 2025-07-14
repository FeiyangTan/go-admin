package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func PKCS7Unpad(data []byte) ([]byte, error) {
	l := len(data)
	if l == 0 {
		return nil, errors.New("invalid data")
	}
	pad := int(data[l-1])
	if pad < 1 || pad > aes.BlockSize {
		return nil, errors.New("invalid padding size")
	}
	return data[:l-pad], nil
}

// DecryptWeChatData 用于解密 encryptedData
func DecryptWeChatData(sessionKey, encryptedData, iv string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	cipherData, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	decrypted := make([]byte, len(cipherData))
	mode.CryptBlocks(decrypted, cipherData)

	return PKCS7Unpad(decrypted)
}
