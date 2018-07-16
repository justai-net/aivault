package aivault

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

//ReadFile reads the file provided and returns bytes
func ReadFile(file string) []byte {
	plaintext, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

//OutToFile given the byte value output to file
func OutToFile(plainCipherText []byte, file string) {
	f, err := os.Create(file)
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(plainCipherText))
	if err != nil {
		panic(err.Error())
	}
}
