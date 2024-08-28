package main

import "fmt"

type RsaEncrypter struct {}

func (rsaEnc RsaEncrypter) Encrypt(plaintext []byte) {
	fmt.Printf("encrypting [%s] with rsa-public-key...\n", string(plaintext))
}

type RsaDecrypter struct {}

func (rsaDec RsaDecrypter) Decrypt(ciphertext []byte) {
	fmt.Printf("decrypting [%02x] with rsa-private-key...\n", ciphertext)
}