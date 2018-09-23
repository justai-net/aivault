package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

//CreateHash creates the hash of a given string and returns the hash
func CreateHash(key string) []byte {
	//hasher := md5.New()
	//hasher.Write([]byte(key))
	//return hex.EncodeToString(hasher.Sum(nil))
	hasher := sha256.Sum256([]byte(key))
	return hasher[:]
}

//Encrypt encrypts the byte file with AES and given passphrase
func Encrypt(data []byte) []byte {
	passphrase := Credentials(true)
	block, _ := aes.NewCipher([]byte(CreateHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	fmt.Println("Encryption was successful")
	return ciphertext
}

//Decrypt decrypts the byte file with AES and given passphrase
func Decrypt(data []byte) []byte {
	passphrase := Credentials(false)
	key := []byte(CreateHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic("Incorrect passphrase, please try")
	}
	gcm, err := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Decryption was successful")
	return plaintext
}

//ViewDecrypted print out decrypted values to terminal
func ViewDecrypted(data []byte) []byte {
	passphrase := Credentials(false)
	key := []byte(CreateHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic("Incorrect passphrase, please try")
	}
	gcm, err := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
