package main

type AesFactory struct {}

func (aesFac AesFactory) createEncrypter() Encrypter {
	return AesEncrypter{}
}

func (aesFac AesFactory) createDecrypter() Decrypter {
	return AesDecrypter{}
}
