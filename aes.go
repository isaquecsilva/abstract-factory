package main

import "fmt"

type AesEncrypter struct {}

func (aesEnc AesEncrypter) Encrypt(plaintext []byte) {
	fmt.Printf("encrypting [%s] with AES...\n", string(plaintext))
}

type AesDecrypter struct {}

func (aesDec AesDecrypter) Decrypt(ciphertext []byte) {
	fmt.Printf("decrypting [%02x] with AES\n", ciphertext)
}