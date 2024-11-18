package main

import (
	"fmt"
	gomail "gopkg.in/mail.v2"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	dirMap := make(map[string][]string) // Maps directory name to names of files that haven't been chosen

	if len(os.Args) == 1 {
		fmt.Println("Must include at least one directory path to choose from.")
		os.Exit(1)
	}

	// Add the specified directories into the map with empty lists as their values
	for _, s := range os.Args[1:] {
		dirMap[s] = []string{}
	}

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
		for dirName, dirList := range dirMap {
			if len(dirList) == 0 {
				dirMap[dirName] = repopulate(dirName)
			}
		}

		message := mdToHTML(smartSelect(dirMap))
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

func repopulate(dir string) []string { // Fills and returns slice containing all files in directory specified
	res := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
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

func smartSelect(dirMap map[string][]string) []string {
	result := []string{} // store n random filepaths in here (n = # specified in command line args)

	for dirName, dirList := range dirMap {
		if len(dirList) == 0 {
			// This should never happen because smartSelect is always called after potential repopulates
			fmt.Printf("Failed on iteration for %s\n", dirName)
			os.Exit(1)
		}

		randNum := rand.Intn(len(dirList)) // Generate random number between 0 and num of files
		result = append(result, chooseFile(dirName, randNum))
	}

	for _, item := range result {
		delete(dirMap, item)
	}

	fmt.Println("Chosen files: ", result)

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
