package main

import (
	"fmt"
	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	dirMap             = make(map[string][]string) // Maps directory name to names of files that haven't been chosen
	ImgDir      string = ""
	webContent  string = ""
	contentLock sync.RWMutex
)

func main() {
	// Command-line argument error handling
	if len(os.Args) <= 3 {
		fmt.Println("Must include at least one directory path to choose from as well as an image directory path.")
		os.Exit(1)
	}

	// Add the specified directories into the map with empty lists as their values (don't include last item)
	for _, s := range os.Args[1 : len(os.Args)-1] {
		dirMap[s] = []string{}
	}

	ImgDir = os.Args[len(os.Args)-1]

	fmt.Printf("Directory arguments: %s\n", dirMap)
	fmt.Printf("Image directory: %s\n\n", ImgDir)

	// Concurrent goroutine for updating webpage dynamically
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()

		fmt.Println("\nSelecting new files...")
		updateContent()
		for range ticker.C {
			fmt.Println("\nSelecting new files...")
			updateContent()
		}
	}()

	mux := http.NewServeMux()

	// Handle requests to get the root
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got / request")
		contentLock.RLock()
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, webContent)
		contentLock.RUnlock()
	})

	mux.HandleFunc("/output/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Requested file: ", r.URL.Path)

		http.ServeFile(w, r, "."+r.URL.Path)
	})

	mux.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got css request")

		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "."+r.URL.Path)
	})

	mux.HandleFunc("/index.js", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got index.js request")

		w.Header().Set("Content-Type", "text/javascript")
		http.ServeFile(w, r, "."+r.URL.Path)
	})

	server := &http.Server{
		Addr:    ":3333",
		Handler: mux,
	}

	fmt.Println("Server is running on http://localhost:3333")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server error: %s\n", err)
		os.Exit(1)
	}
}

func updateContent() {
	// If any of the arrays are empty, then repopulate them and begin a new cycle
	for dirName, dirList := range dirMap {
		if len(dirList) == 0 {
			fmt.Printf("Repopulating %s...\n\n", dirName)
			dirMap[dirName] = repopulate(dirName)
		}
	}

	selectedFiles := smartSelect(dirMap)

	copiedFiles := copyFile(selectedFiles)
	_ = copyImages(copiedFiles)

	message := mdToHTML(copiedFiles)

	contentLock.Lock()
	webContent = string(message)

	//sendEmail(string(message))
	contentLock.Unlock()
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
	var totalSize int64 = 0

	for dirName, dirList := range dirMap {
		if len(dirList) == 0 {
			// This should never happen because smartSelect is always called after potential repopulates
			fmt.Printf("Failed on iteration for %s\n", dirName)
			os.Exit(1)
		}

		randNum := rand.Intn(len(dirList)) // Generate random number between 0 and num of files
		chosen := chooseFile(dirName, randNum)

		info, err := os.Stat(chosen)
		if err != nil {
			fmt.Printf("Problem obtaining the file info of the file: %s\n", chosen)
			continue
		}

		totalSize += info.Size()

		// If the totalSize of the files we wish to add is larger than 20 kB, don't add the additional file
		if len(result) > 1 && totalSize > 20480 { // Maybe change to >= 1
			totalSize -= info.Size()
			break
		}

		result = append(result, chosen)
	}

	for _, item := range result {
		delete(dirMap, item)
	}

	fmt.Printf("Chosen files: %s. Total size: %d.\n", result, totalSize)

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

func sendEmail(message string) {
	err := godotenv.Load(".gitignore/.env")

	if err != nil {
		fmt.Println("Error loading .env file.")
		os.Exit(1)
	}

	from := os.Getenv("FROM")
	mailPass := os.Getenv("PASSWORD")
	to := os.Getenv("TO")
	smtpHost := "smtp.gmail.com"

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Review Notes")

	m.SetBody("text/html", message)

	d := gomail.NewDialer(smtpHost, 587, from, mailPass)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("Email sent successfully!")
}
