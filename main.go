package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	host = "smtp.gmail.com"
	addr = host + ":587"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gomail <recipent_address>")
		os.Exit(1)
	}

	from := os.Getenv("GOMAIL_USER")
	to := os.Args[1:]

	message := []byte{}
	f, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// Read from stdin if not empty else use template for message
	if f.Size() > 0 {
		message, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Open with default editor if stdin is empty
		tempFile := "/tmp/gomail"

		// Overwrite temp file with template file
		template := make([]byte, 0)
		templateFile := os.Getenv("HOME") + "/.gomail_template"
		if _, err := os.Stat(templateFile); err == nil {
			template, err = ioutil.ReadFile(templateFile)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = ioutil.WriteFile(tempFile, template, 0644)
		if err != nil {
			log.Fatal(err)
		}

		// Open temp file with default editor
		editor := os.Getenv("EDITOR")
		cmd := exec.Command(editor, tempFile)
		cmd.Stdin, cmd.Stdout = os.Stdin, os.Stdout
		cmd.Run()

		// Set message from temp file user just edited
		message, err = ioutil.ReadFile(tempFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Authenticate and send mail
	fmt.Printf("Enter password: ")
	tmp, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	password := string(tmp)
	if err != nil {
		log.Fatal(err)
	}
	password = strings.Trim(password, " \t\n")
	auth := smtp.PlainAuth("", from, password, host)
	fmt.Print("Authenticating... ")
	fmt.Println("Success")

	// Send mail
	fmt.Print("Sending mail... ")
	if err := smtp.SendMail(addr, auth, from, to, message); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success")
}
