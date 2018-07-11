package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/justai-net/aivault"
)

func main() {

	fileFlags := flag.NewFlagSet("encrypt", flag.ExitOnError)
	//subDecrypt := flag.NewFlagSet("decrypt", flag.ExitOnError)
	file := fileFlags.String("file", "", "Selected file to be encrypted/decrypted")
	flag.Parse()

	switch os.Args[1] {
	case "encrypt":
		fileFlags.Parse(os.Args[2:])
		fmt.Println(*file)
		ciphertext := aivault.Encrypt(aivault.ReadFile(*file))
		aivault.OutToFile(ciphertext)
	case "decrypt":
		fileFlags.Parse(os.Args[2:])
		plaintext := aivault.Decrypt(aivault.ReadFile(*file))
		aivault.OutToFile(plaintext)
	}
	//fmt.Printf("Password: %s\n", passphrase)
	//fmt.Println(createHash(passphrase))
	aivault.ReadFile()
	//fmt.Println(ciphertext)

}
