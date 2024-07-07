package main

import (
	"fmt"
	gomail "gopkg.in/mail.v2"
	"math/rand"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
	"unicode/utf8"
)

var cpscDir []string
var ssDir []string
var htstDir []string

var cpscPathTo string = ""
var ssPathTo string = ""
var htstPathTo string = ""

var images []string

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

		message, messageMD := mdToHTML(smartSelect())
		images := extractPhotos(messageMD)
		messageAsString := string(message)

		m.SetBody("text/html", messageAsString)

		// add attachments
		var imagePath []string
		var imageDir string = ""
		err := filepath.Walk(imageDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println("Error adding attachments: ", err)
				return err
			}
			if !info.IsDir() {
				// if images were created/modified after a certain time, do...
				if slices.Contains(images, info.Name()) {
					imagePath = append(imagePath, path)
				}
			}
			return nil
		})

		fmt.Println("Paths to the images: ", imagePath)

		if err != nil {
			fmt.Println(err)
		}

		for _, item := range imagePath {
			m.Attach(item)
		}

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

	// This should never happen because smartSelect is always called after potential repopulates
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

func extractPhotos(original []byte) []string {
	var images []string

	var excalamationRune bool = false
	var leftSquareRune bool = false
	var rightSquareRune bool = false
	var leftParenRune bool = false
	var rightParenRune bool = false
	var consecCount int = 0

	var toWrite strings.Builder

	// 33 is ! ; 91 is [
	for len(original) > 0 {
		rune, size := utf8.DecodeRune(original)

		// CHARACTERS HAVE TO BE IN SUCCESSION somewhat
		if !excalamationRune && !leftSquareRune && !rightSquareRune && !leftParenRune && rune == 33 {
			excalamationRune = true
			consecCount++
		} else if excalamationRune && !leftSquareRune && !rightSquareRune && !leftParenRune && rune == 91 && consecCount == 1 {
			leftSquareRune = true
			consecCount++
		} else if excalamationRune && leftSquareRune && !rightSquareRune && !leftParenRune && rune == 93 && consecCount == 2 {
			rightSquareRune = true
			consecCount++
		} else if excalamationRune && leftSquareRune && rightSquareRune && !leftParenRune && rune == 40 && consecCount == 3 {
			leftParenRune = true
			consecCount++
		} else if excalamationRune && leftSquareRune && rightSquareRune && leftParenRune && !rightParenRune && rune != 41 {
			if rune != 60 && rune != 62 {
				toWrite.WriteRune(rune)
			}
		} else if excalamationRune && leftSquareRune && rightSquareRune && leftParenRune && !rightParenRune && rune == 41 {
			rightParenRune = true
			image := toWrite.String()
			images = append(images, image)
			toWrite.Reset()

			excalamationRune = false
			leftSquareRune = false
			rightSquareRune = false
			leftParenRune = false
			rightParenRune = false
			consecCount = 0
		}

		original = original[size:]
	}

	return images
}
