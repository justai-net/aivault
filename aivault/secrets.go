package aivault

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

//StringInSlice comment
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

//ReadFile reads the file provided and returns bytes
func ReadFile(file string) []byte {
	plaintext, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

//CreateHash creates the hash of a given key and returns the hash
func CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

//Encrypt encrypts the byte file with AES and given passphrase
func Encrypt(data []byte) []byte {
	passphrase := Credentials()
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
	return ciphertext
}

//Decrypt decrypts the byte file with AES and given passphrase
func Decrypt(data []byte) []byte {
	passphrase := Credentials()
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

//OutToFile given the byte value output to file
func OutToFile(plainCipherText []byte) {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(plainCipherText))
	if err != nil {
		panic(err.Error())
	}
}

//Credentials prompts the user for a password
func Credentials() (password string) {
	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err.Error())
	}
	password = string(bytePassword)

	return
}
