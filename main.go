package main

import (
	"log"
	"errors"
)

type Loader struct {
	facs map[string]EncryptionFactory
}

func (l *Loader) Add(algoName string, fac EncryptionFactory) error {
	if _, ok := l.facs[algoName]; ok {
		return errors.New("already defined factory")
	}

	l.facs[algoName] = fac
	return nil
}

func (l *Loader) Load(algoName string) (EncryptionFactory, error) {
	fac, ok := l.facs[algoName]

	if !ok {
		return nil, errors.New("factory not found")
	}

	return fac, nil	
}

func main() {
	loader := &Loader{facs: make(map[string]EncryptionFactory)}
	loader.Add("aes", AesFactory{})
	loader.Add("rsa", RsaFactory{})

	// encrypting with rsa
	clientCode("hello, world!", loader, "rsa")
	clientCode("hello, world!", loader, "aes")
}

func clientCode(msg string, loader *Loader, algoName string) {
	fac, err := loader.Load(algoName)

	if err != nil {
		log.Fatal(err)
	}

	encrypter := fac.createEncrypter()
	decrypter := fac.createDecrypter()

	encrypter.Encrypt([]byte(msg))
	decrypter.Decrypt([]byte(msg))
}