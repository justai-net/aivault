package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	aivault "github.com/justai-net/aivault/secrets"
)

func main() {

	flag.Usage = func() {
		fmt.Println("Usage: aivault [encrypt|decrypt|view|copy] [file]")
	}
	flag.Parse()
	if len(os.Args) <= 2 {
		flag.Usage()
		os.Exit(1)
	}

	file := os.Args[2]
	arch := runtime.GOOS

	switch os.Args[1] {
	case "encrypt":
		//fileFlags.Parse(os.Args[2:])
		l := log.New(os.Stderr, "", 0)
		os.Stderr.WriteString("Encrypting file " + file + " with AES-256")
		l.Println()
		ciphertext := aivault.Encrypt(aivault.ReadFile(file))
		aivault.OutToFile(ciphertext, file)

	case "decrypt":
		//fileFlags.Parse(os.Args[2:])
		plaintext := aivault.Decrypt(aivault.ReadFile(file))
		aivault.OutToFile(plaintext, file)

	case "view":
		plaintext := aivault.ViewDecrypted(aivault.ReadFile(file))
		if len(os.Args) <= 3 {
			accounts := aivault.GetAllAccounts(plaintext)
			os.Stderr.WriteString("Available Accounts")
			fmt.Println()
			for i := range accounts {
				fmt.Println(accounts[i])
			}
		} else {
			var s []byte
			s = aivault.GetAccount(plaintext, os.Args[3])
			fmt.Println(string(s))
		}

	case "copy":
		if len(os.Args) <= 3 {
			os.Stderr.WriteString("Enter Selected username to get password: ")
			os.Exit(1)
		}
		plaintext := aivault.ViewDecrypted(aivault.ReadFile(file))
		aivault.ToClipboard(aivault.GetPassword(plaintext, os.Args[3]), arch)
	}
}
