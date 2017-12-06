package aesECB

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
)

// 加密
func Encrypt(plaintext string, key string) []byte {
	cipher, err := aes.NewCipher([]byte(key[:aes.BlockSize]))
	if err != nil {
		panic(err.Error())
	}

	padText := PKCS7Pad([]byte(plaintext))
	if len(padText)%aes.BlockSize != 0 {
		panic("Need a multiple of the blocksize 16")
	}

	ciphertext := make([]byte, 0)
	text := make([]byte, 16)
	for len(padText) > 0 {
		// 每次运算一个block
		cipher.Encrypt(text, padText)
		padText = padText[aes.BlockSize:]
		ciphertext = append(ciphertext, text...)
	}
	return ciphertext
}

// 解密
func Decrypt(ciphertext string, key string) []byte {
	cipher, err := aes.NewCipher([]byte(key[:aes.BlockSize]))
	if err != nil {
		panic(err.Error())
	}
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("Need a multiple of the blocksize 16")
	}

	plaintext := make([]byte, 0)
	text := make([]byte, 16)
	byteText, _ := hex.DecodeString(ciphertext)
	for len(byteText) > 0 {
		cipher.Decrypt(text, byteText)
		byteText = byteText[aes.BlockSize:]
		plaintext = append(plaintext, text...)
	}
	plaintext = PKCS7UPad(plaintext)
	return plaintext
}

// Padding补全
func PKCS7Pad(data []byte) []byte {
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

//
func PKCS7UPad(data []byte) []byte {
	padLength := int(data[len(data)-1])
	return data[:len(data)-padLength]
}
