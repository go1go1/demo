package main

import (
	"demo/lib"
	"fmt"
	// "crypto/md5"

	"encoding/base64"
)

func main() {
	testAesECB()
}

func testAesECB() {
	fmt.Println("testAesECB ---------------------------------")
	txt := "201711270101000000001"
	key := "4uoRHo0TC62DyLPh7QYlWA=="

	heKey, _ := base64.StdEncoding.DecodeString(key)
	cipherText := lib.Encrypt(lib.PKCS7Pad([]byte(txt)), string(heKey))
	fmt.Printf("ciphertext: %x \n", cipherText)
	fmt.Printf("base64 ciphertext: %s \n", base64.StdEncoding.EncodeToString(cipherText))
	decryptText := lib.Decrypt(cipherText, string(heKey))
	plainText := lib.PKCS7UPad([]byte(decryptText))
	fmt.Printf("plaintext: %s \n", plainText)

}
