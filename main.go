package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

var usage = `A command-line interface for secrets management.

Usage:
    aivault [command] [flags] file_name

Available Commands:
    encrypt     Encrypt file with a passphrase.
    decrypt     Decrypt file with the original passphrase.
    view        View account names or username if supplied.
    copy        Copy password into clipboard.
    add         Add a new entry into the file

Flags:
    -h, --help  help for this command

Use "[command] --help"  for more information about a command.
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, string(usage))
	}

	flag.Parse()
	// Subcommands
	encryptCommand := flag.NewFlagSet("encrypt", flag.ExitOnError)
	decryptCommand := flag.NewFlagSet("decrypt", flag.ExitOnError)
	viewCommand := flag.NewFlagSet("view", flag.ExitOnError)
	copyCommand := flag.NewFlagSet("copy", flag.ExitOnError)
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)

	// Add subcommand flags pointers
	addCommand.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: aivault add [flags] file_name \n")
		addCommand.PrintDefaults()
	}
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
		fmt.Fprintf(os.Stderr, "Encrypting file "+file+" with AES-256")
		fmt.Fprintf(os.Stderr, "\n")
		ciphertext := Encrypt(ReadFile(file))
		OutToFile(ciphertext, file)
	case "decrypt":
		decryptCommand.Parse(os.Args[2:])
		plaintext := Decrypt(ReadFile(file))
		OutToFile(plaintext, file)
	case "view":
		viewCommand.Parse(os.Args[2:])
		plaintext := ViewDecrypted(ReadFile(file))
		if len(os.Args) <= 3 {
			accounts := GetAllAccounts(plaintext)
			fmt.Fprintf(os.Stderr, "Available Accounts")
			fmt.Println()
			for i := range accounts {
				fmt.Println(accounts[i])
			}
		} else {
			GetAccount(plaintext, os.Args[3])
		}
		fmt.Fprintf(os.Stderr, "\nUsage to copy password to clipboard:\naivault copy file_name [account_name]\n")

	case "copy":
		copyCommand.Parse(os.Args[2:])
		if len(os.Args) <= 3 {
			os.Stderr.WriteString("Enter Selected username to get password: ")
			os.Exit(1)
		}
		plaintext := ViewDecrypted(ReadFile(file))
		ToClipboard(GetPassword(plaintext, os.Args[3]), arch)
	case "add":
		addCommand.Parse(os.Args[2:])
		fmt.Println(*addCommandaccount)
		fmt.Println(*addCommandusername)
	}
}
