package aivault

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

//Credentials prompts the user for a password with confirmation
func Credentials(confirm bool) (password string) {
	fmt.Fprintf(os.Stderr, "Password: ") //To ignore pipes
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println() // Stderr new line, ignore pipes
	if err != nil {
		panic(err.Error())
	}
	if confirm == true {
		fmt.Fprintf(os.Stderr, "Confirm Password: ")
		confirmPassword, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println() // Stderr new line, ignore pipes
		if err != nil {
			panic(err.Error())
		}
		if string(bytePassword) != string(confirmPassword) {
			fmt.Fprintf(os.Stderr, "Passwords do not match")
			fmt.Println()
			os.Exit(1)
		}
	}
	password = string(bytePassword)
	return
}
