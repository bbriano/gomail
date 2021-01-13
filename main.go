package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"os/exec"
)

const (
	templateFile = "~/.gomail_template"
)

func main() {
	from := os.Getenv("GOMAIL_USER")
	password := os.Getenv("GOMAIL_PASS")
	to := os.Args[1:]
	host := "smtp.gmail.com"
	addr := host + ":587"
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
		template, err := ioutil.ReadFile(templateFile)
		if err != nil {
			log.Fatal(err)
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

	// Send to self if not specify receiver
	if len(to) <= 0 {
		to = []string{from}
	}

	// Authenticate and send mail
	auth := smtp.PlainAuth("", from, password, host)
	if err := smtp.SendMail(addr, auth, from, to, message); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mail Sent!")
}
