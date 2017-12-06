package lib

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	key := "thisiskeyandlen>16"
	ciphertext := Encrypt(PKCS7Pad([]byte("asdasdasdasdasd")), key)
	fmt.Printf("ciphertext: %x \n", ciphertext)
	plaintext := Decrypt(ciphertext, key)
	fmt.Printf("plaintext: %s \n", PKCS7UPad([]byte(plaintext)))
}
