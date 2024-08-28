package main

type Encrypter interface {
	Encrypt([]byte)
}

type Decrypter interface {
	Decrypt([]byte)
}