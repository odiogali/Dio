package main

import (
	"fmt"
	gomail "gopkg.in/mail.v2"
	"math/rand"
	"os"
	"path/filepath"
	// "strings"
	"time"
)

var cpscDir []string
var ssDir []string
var htstDir []string

var cpscPathTo string = ""
var ssPathTo string = ""
var htstPathTo string = ""

func main() {
	for {
		from := ""
		mailPass := ""
		to := ""
		smtpHost := "smtp.gmail.com"

		m := gomail.NewMessage()

		m.SetHeader("From", from)
		m.SetHeader("To", to)
		m.SetHeader("Subject", "Review Notes")

		// If any of the arrays are empty, then repopulate them and begin a new cycle
		if cpscDir == nil {
			cpscDir = repopulate(cpscPathTo)
		}

		if ssDir == nil {
			ssDir = repopulate(ssPathTo)
		}

		if htstDir == nil {
			htstDir = repopulate(htstPathTo)
		}

		message := mdToHTML(smartSelect())
		messageAsString := string(message)

		m.SetBody("text/html", messageAsString)

		d := gomail.NewDialer(smtpHost, 587, from, mailPass)

		if err := d.DialAndSend(m); err != nil {
			fmt.Println(err)
			panic(err)
		}

		fmt.Println("Email sent successfully!")

		time.Sleep(24 * time.Hour)
	}
}

func repopulate(dir string) []string {
	var stringPath string = dir
	res := []string{}

	err := filepath.Walk(stringPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() {
			res = append(res, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return res
}

func smartSelect() []string {
	result := []string{} // store 3 random filepaths in here

	if len(cpscDir) == 0 || len(ssDir) == 0 || len(htstDir) == 0 {
		fmt.Println("CPSC: ", len(cpscDir), " SS: ", len(ssDir), " HTST: ", len(htstDir))
		return result
	}

	cpscRand := rand.Intn(len(cpscDir)) // Generate random number between 0 and num of files
	ssRand := rand.Intn(len(ssDir))
	htstRand := rand.Intn(len(htstDir))

	result = append(result, chooseFile(cpscPathTo, cpscRand))
	result = append(result, chooseFile(ssPathTo, ssRand))
	result = append(result, chooseFile(htstPathTo, htstRand))

	fmt.Println(result)

	return result
}

func chooseFile(direcName string, randInt int) string {
	var counter int = 0
	result := ""

	err := filepath.Walk(direcName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Choose File error: ", err)
			return err
		}
		if !info.IsDir() {
			if counter == randInt {
				result = path
			}
			counter++
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return result
}
