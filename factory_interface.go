package main

type EncryptionFactory interface {
	createEncrypter() Encrypter
	createDecrypter() Decrypter
}