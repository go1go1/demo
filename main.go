package main

import (
	"crypto/aes"
	"demo/aesCBC"
	"demo/aesECB"
	"demo/desCBC"
	"demo/desECB"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func main() {
	// testAesECB()
	// testAesCBC()
	// testDesCBC()
	testDesECB()
}

func testAesECB() {
	fmt.Println("testAesECB ---------------------------------")
	txt := "201711270101000000001"
	key := "4uoRHo0TC62DyLPh7QYlWA=="

	heKey, _ := base64.StdEncoding.DecodeString(key)
	// byteData := aesECB.Encrypt(aesECB.PKCS7Pad([]byte(txt), 128), string(heKey))
	byteData := aesECB.Encrypt(aesECB.PKCS7Pad([]byte(txt), aes.BlockSize), string(heKey))
	cipherText := hex.EncodeToString(byteData)
	fmt.Println(cipherText)
	fmt.Printf("base64 ciphertext: %s \n", base64.StdEncoding.EncodeToString(byteData))
	plainText := aesECB.PKCS7UnPad(aesECB.Decrypt(cipherText, string(heKey)))
	fmt.Printf("plaintext: %s \n", plainText)

}

func testAesCBC() {
	fmt.Println("testAesCBC ---------------------------------")
	txt := "201711270101000000001"
	key := "4uoRHo0TC62DyLPh7QYlWA=="

	heKey, _ := base64.StdEncoding.DecodeString(key)
	byteData, _ := aesCBC.Encrypt([]byte(txt), heKey)
	cipherText := hex.EncodeToString(byteData)
	fmt.Println(cipherText)
	fmt.Printf("base64 ciphertext: %s \n", base64.StdEncoding.EncodeToString(byteData))
	plainText, _ := aesCBC.Decrypt(byteData, heKey)
	fmt.Printf("plaintext: %s \n", plainText)
}

func testJSON() string {
	fmt.Println("testJSON ---------------------------------")
	token := "7d2758efad43f37ff1957018429c3946"
	type Info struct {
		Token       string `json:"token"`
		Name        string `json:"name"`
		Cardno      string `json:"cardno"`
		Mobile      string `json:"mobile"`
		CompanyName string `json:"companyName"`
		Salary      string `json:"salary"`
	}
	info := Info{
		Token:       token,
		Name:        "木木",
		Cardno:      "110000000000000000",
		Mobile:      "",
		CompanyName: "北京xxxx有限公司",
		Salary:      "1000",
	}

	txt, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	res := string(txt)
	fmt.Println(res)
	return res
}

func testDesCBC() {
	fmt.Println("testDesCBC ---------------------------------")
	txt := "201711270101000000001"
	key := "12345678"

	result, _ := desCBC.DesEncrypt([]byte(txt), []byte(key))
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, _ := desCBC.DesDecrypt(result, []byte(key))
	fmt.Println(string(origData))
}

func testDesECB() {
	fmt.Println("testDesECB ---------------------------------")
	txt := testJSON()
	key := "788f2a_A851668_Q"
	result, _ := desECB.DesEncrypt([]byte(txt), []byte(key))
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, _ := desECB.DesDecrypt(result, []byte(key))
	fmt.Println(string(origData))
}
