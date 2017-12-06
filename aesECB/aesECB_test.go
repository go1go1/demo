package aesECB

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	key := "thisiskeyandlen>16"
	text := "asdasdasdasdasd"
	ciphertext := Encrypt(text, key)
	fmt.Printf("ciphertext: %x \n", ciphertext)
	plaintext := Decrypt(ciphertext, key)
	fmt.Printf("plaintext: %s \n", PKCS7UPad([]byte(plaintext)))
}
