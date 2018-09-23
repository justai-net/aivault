package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
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

//ToClipboard copies the byte into your clipboard as a string
func ToClipboard(output []byte, arch string) {
	var copyCmd *exec.Cmd

	// Mac "OS"
	if arch == "darwin" {
		copyCmd = exec.Command("pbcopy")
	}
	// Linux
	if arch == "linux" {
		copyCmd = exec.Command("xclip", "-selection", "c")
	}

	in, err := copyCmd.StdinPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := copyCmd.Start(); err != nil {
		log.Fatal(err)
	}

	if _, err := in.Write([]byte(output)); err != nil {
		log.Fatal(err)
	}

	if err := in.Close(); err != nil {
		log.Fatal(err)
	}

	copyCmd.Wait()
}

//GetPassword returns only password value
func GetPassword(output []byte, account string) (passValue []byte) {
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		a, u, p := s[0], s[1], s[2]
		if a == account {
			fmt.Println("Account: " + a)
			fmt.Println("Username: " + u)
			passValue = []byte(p)
		}
	}
	return
}

//GetAccount returns accountname and username
func GetAccount(output []byte, account string) (accountValue []byte, userValue []byte) {
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		a, u, _ := s[0], s[1], s[2]
		if a == account {
			accountValue = []byte(a)
			userValue = []byte(u)
			fmt.Fprintf(os.Stderr, "Found Account!\n")
			fmt.Fprintf(os.Stderr, "Account Name: "+string(accountValue)+"\n")
			fmt.Fprintf(os.Stderr, "Username: "+string(userValue)+"\n")
			fmt.Fprintf(os.Stderr, "Password: **********\n")
			return
		}
	}
	fmt.Fprintf(os.Stderr, "Account %s not found in this file.\n", string(account))
	os.Exit(1)
	return
}

//GetAllAccounts returns all Accounts
func GetAllAccounts(output []byte) (user []string) {
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		user = append(user, s[0])
	}
	return
}
