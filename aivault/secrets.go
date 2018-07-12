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
	"log"
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

//Credentials prompts the user for a password with confirmation
func Credentials(confirm bool) (password string) {
	os.Stderr.WriteString("Password: ") //To ignore pipes
	l := log.New(os.Stderr, "", 0)
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	l.Println() // Stderr new line, ignore pipes
	if err != nil {
		panic(err.Error())
	}
	if confirm == true {
		os.Stderr.WriteString("Confirm Password: ")
		confirmPassword, err := terminal.ReadPassword(int(syscall.Stdin))
		l.Println() // // Stderr new line, ignore pipes
		if err != nil {
			panic(err.Error())
		}
		if string(bytePassword) != string(confirmPassword) {
			fmt.Println("Passwords do not match")
			os.Exit(1)
		}
	}
	password = string(bytePassword)

	return
}

//ViewDecrypted print out decrypted values to terminal
func ViewDecrypted(data []byte) {
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
	os.Stderr.WriteString(string(plaintext))
}
