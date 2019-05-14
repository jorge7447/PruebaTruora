package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		os.Exit(1)
	}
	return privkey, &privkey.PublicKey
}

//Funcion para convertir la llave publica en bytes
func PublicKeyToBytes(publicKey *rsa.PublicKey) []byte {

	/*pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		log.Error(err)
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes*/

	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	public := pem.EncodeToMemory(block)

	return public
}

//Funcion para convertir la llave privada en bytes
func PrivateKeyToBytes(privateKey *rsa.PrivateKey, password string) []byte {
	/*privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	return privBytes*/
	// Convert it to pem
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	var err error
	// Encrypt the pem
	block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, []byte(password), x509.PEMCipherAES256)
	if err != nil {
		os.Exit(1)
	}

	private := pem.EncodeToMemory(block)

	return private
}

//Funcion para convertir los bytes en una llave publica
func BytesToPublicKey(public []byte) *rsa.PublicKey {

	/*block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			log.Error(err)
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		log.Error(err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		log.Error("not ok")
	}
	return key*/

	block, _ := pem.Decode(public) //[]byte(keyData.Public)
	result, _ := x509.ParsePKCS1PublicKey(block.Bytes)
	//fmt.Println(parseResult)
	return result
}

func BytesToPrivateKey(private []byte, password string) *rsa.PrivateKey {

	/*block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			log.Error(err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		log.Error(err)
	}
	return key*/

	var err error
	block, _ := pem.Decode(private) //[]byte(keyData.Public)

	encrypted := x509.IsEncryptedPEMBlock(block)
	bytes := block.Bytes

	if encrypted {
		bytes, err = x509.DecryptPEMBlock(block, []byte(password))
		if err != nil {
			log.Panic(err)
			fmt.Println(err)
		}
	}

	/*
		result, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		return result*/
	result, err := x509.ParsePKCS1PrivateKey(bytes)
	if err != nil {
		log.Panic(err)
		fmt.Println(err)
	}

	return result
}

func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		log.Panic(err)
	}
	return ciphertext
}

func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) []byte {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		log.Panic(err.Error)
	}
	return plaintext
}
