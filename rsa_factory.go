package main

type RsaFactory struct {}

func (rsaFac RsaFactory) createEncrypter() Encrypter {
	return RsaEncrypter{}
}

func (rsaFac RsaFactory) createDecrypter() Decrypter {
	return RsaDecrypter{}
}