package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/justai-net/aivault/secrets"
)

func main() {

	flag.Usage = func() {
		fmt.Println("Usage: aivault [encrypt|decrypt|view] [file]")
	}
	flag.Parse()
	file := os.Args[2]

	switch os.Args[1] {
	case "encrypt":
		//fileFlags.Parse(os.Args[2:])
		fmt.Println(file)
		ciphertext := aivault.Encrypt(aivault.ReadFile(file))
		aivault.OutToFile(ciphertext, file)
	case "decrypt":
		//fileFlags.Parse(os.Args[2:])
		plaintext := aivault.Decrypt(aivault.ReadFile(file))
		aivault.OutToFile(plaintext, file)
	case "view":
		plaintext := aivault.ViewDecrypted(aivault.ReadFile(file))
		s := string(plaintext)
		fmt.Print(s)
	}

}
