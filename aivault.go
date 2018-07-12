package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/justai-net/aivault"
)

func main() {

	flag.Usage = func() {
		fmt.Println("Usage: aivault [encrypt|decrypt|view] [file]")
	}
	flag.Parse()
	//var file string
	file := os.Args[2]

	switch os.Args[1] {
	case "encrypt":
		//fileFlags.Parse(os.Args[2:])
		fmt.Println(file)
		ciphertext := aivault.Encrypt(aivault.ReadFile(file))
		aivault.OutToFile(ciphertext)
	case "decrypt":
		//fileFlags.Parse(os.Args[2:])
		plaintext := aivault.Decrypt(aivault.ReadFile(file))
		aivault.OutToFile(plaintext)
	case "view":
		aivault.ViewDecrypted(aivault.ReadFile(file))
	}

}
