package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"log"
)

func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Panic(err)
	}

	if err := privateKey.Validate(); err != nil {
		log.Panic(err)
	}

	if privateKey.D.Cmp(privateKey.N) > 0 {
		log.Panic("Exponente privado demasiado grande")
	}

	return privateKey, &privateKey.PublicKey
}

//Funcion para convertir la llave publica en bytes
func PublicKeyToBytes(publicKey *rsa.PublicKey) []byte {
	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	public := pem.EncodeToMemory(block)

	return public
}

//Funcion para convertir la llave privada en bytes
func PrivateKeyToBytes(privateKey *rsa.PrivateKey, password string) []byte {
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	var err error
	block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, []byte(password), x509.PEMCipherAES256)
	if err != nil {
		log.Panic(err)
	}

	private := pem.EncodeToMemory(block)

	return private
}

//Funcion para convertir los bytes en una llave publica
func BytesToPublicKey(public []byte) *rsa.PublicKey {

	block, _ := pem.Decode(public)
	result, _ := x509.ParsePKCS1PublicKey(block.Bytes)
	return result
}

//Funcion para convertir los bytes en una llave privada
func BytesToPrivateKey(private []byte, password string) *rsa.PrivateKey {

	var err error
	block, _ := pem.Decode(private)

	encrypted := x509.IsEncryptedPEMBlock(block)
	bytes := block.Bytes

	if encrypted {
		bytes, err = x509.DecryptPEMBlock(block, []byte(password))
		if err != nil {
			log.Panic(err)
		}
	}

	result, err := x509.ParsePKCS1PrivateKey(bytes)
	if err != nil {
		log.Panic(err)
	}

	return result
}

//Funcion para encriptar con una llave publica
func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	ciphertext, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, pub, msg, nil)
	if err != nil {
		log.Panic(err)
	}
	return ciphertext
}

//Funcion para desencriptar con una llave privada
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) []byte {
	plaintext, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, priv, ciphertext, nil)
	if err != nil {
		log.Panic(err.Error)
	}
	return plaintext
}
