package main

import (
	"demo/aesCBC"
	"demo/aesECB"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	testAesECB()
	testDesCBC()
}

func testAesECB() {
	fmt.Println("testAesECB ---------------------------------")
	txt := "201711270101000000001"
	key := "4uoRHo0TC62DyLPh7QYlWA=="

	heKey, _ := base64.StdEncoding.DecodeString(key)
	byteData := aesECB.Encrypt(txt, string(heKey))
	cipherText := hex.EncodeToString(byteData)
	fmt.Println(cipherText)
	fmt.Printf("base64 ciphertext: %s \n", base64.StdEncoding.EncodeToString(byteData))
	plainText := aesECB.Decrypt(cipherText, string(heKey))
	fmt.Printf("plaintext: %s \n", plainText)

}

func testDesCBC() {
	fmt.Println("testAesECB ---------------------------------")
	txt := "201711270101000000001"
	key := "4uoRHo0TC62DyLPh7QYlWA=="

	heKey, _ := base64.StdEncoding.DecodeString(key)
	byteKey := []byte(heKey)
	byteData, _ := aesCBC.AesEncrypt([]byte(txt), byteKey)
	cipherText := hex.EncodeToString(byteData)
	fmt.Println(cipherText)
	fmt.Printf("base64 ciphertext: %s \n", base64.StdEncoding.EncodeToString(byteData))
	plainText, _ := aesCBC.AesDecrypt(byteData, byteKey)
	fmt.Printf("plaintext: %s \n", plainText)
}
