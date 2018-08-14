package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	aivault "github.com/justai-net/aivault/secrets"
)

func main() {
	b, err := ioutil.ReadFile("usage.txt")
	if err != nil {
		fmt.Print(err)
	}
	flag.Usage = func() {
		fmt.Println(string(b))
	}

	flag.Parse()
	// Subcommands
	encryptCommand := flag.NewFlagSet("encrypt", flag.ExitOnError)
	decryptCommand := flag.NewFlagSet("decrypt", flag.ExitOnError)
	viewCommand := flag.NewFlagSet("view", flag.ExitOnError)
	copyCommand := flag.NewFlagSet("copy", flag.ExitOnError)
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)

	// Add subcommand flags pointers
	addCommandaccount := addCommand.String("accountname", "", "Name for the account")
	addCommandusername := addCommand.String("username", "", "Username for the account")

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}
	file := os.Args[2]
	arch := runtime.GOOS

	switch os.Args[1] {
	case "encrypt":
		encryptCommand.Parse(os.Args[2:])
		//fileFlags.Parse(os.Args[2:])
		l := log.New(os.Stderr, "", 0)
		os.Stderr.WriteString("Encrypting file " + file + " with AES-256")
		l.Println()
		ciphertext := aivault.Encrypt(aivault.ReadFile(file))
		aivault.OutToFile(ciphertext, file)

	case "decrypt":
		decryptCommand.Parse(os.Args[2:])
		//fileFlags.Parse(os.Args[2:])
		plaintext := aivault.Decrypt(aivault.ReadFile(file))
		aivault.OutToFile(plaintext, file)

	case "view":
		viewCommand.Parse(os.Args[2:])
		plaintext := aivault.ViewDecrypted(aivault.ReadFile(file))
<<<<<<< HEAD
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
=======
		s := string(plaintext)
		fmt.Print(s)
	}
>>>>>>> c001eda60e060b4d62dff9157a12c1befd2b2812

	case "copy":
		copyCommand.Parse(os.Args[2:])
		if len(os.Args) <= 3 {
			os.Stderr.WriteString("Enter Selected username to get password: ")
			os.Exit(1)
		}
		plaintext := aivault.ViewDecrypted(aivault.ReadFile(file))
		aivault.ToClipboard(aivault.GetPassword(plaintext, os.Args[3]), arch)
	case "add":
		addCommand.Parse(os.Args[2:])
		fmt.Println(*addCommandaccount)
		fmt.Println(*addCommandusername)
	}
}
