package aivault

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

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
