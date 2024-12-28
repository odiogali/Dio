package main

import (
	"fmt"
	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
	"io"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var (
	dirMap             = make(map[string][]string) // Maps directory name to names of files that haven't been chosen
	imgDir      string = ""
	webContent  string = ""
	contentLock sync.RWMutex
	PORT        string = ":3333"
)

func main() {
	// Command-line argument error handling
	if len(os.Args) <= 3 {
		fmt.Println("Must include at least one directory path to choose from as well as an image directory path.")
		os.Exit(1)
	}

	// Add the specified directories into the map with empty lists as their values (don't include last item)
	// All directories but the last are note directories - have markdown notes, not images
	for _, s := range os.Args[1 : len(os.Args)-1] {
		dirMap[s] = []string{}
	}
	// Final argument - which should be directory path - is the image directory
	imgDir = os.Args[len(os.Args)-1]
	// Ensure that user is aware of the above stipulations
	fmt.Printf("Directory arguments: %s\n", dirMap)
	fmt.Printf("Image directory: %s\n\n", imgDir)

	// Concurrent goroutine for updating webpage dynamically
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		fmt.Println("\nSelecting new files...")
		cleanRepo()
		updateContent()
		for range ticker.C {
			fmt.Println("\nSelecting new files...")
			cleanRepo()
			updateContent()
		}
	}()

	mux := http.NewServeMux()

	// Handle requests to get the root
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		contentLock.RLock()
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, webContent)
		contentLock.RUnlock()
	})

	mux.HandleFunc("/output/images/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "."+r.URL.Path)
	})

	mux.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "."+r.URL.Path)
	})

	server := &http.Server{
		Addr:    PORT,
		Handler: mux,
	}

	fmt.Println("Server is running on http://localhost", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(fmt.Errorf("Error occured: %v\n", err))
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

	selectedFiles := smartSelect() // Select files from note directories

	copiedFiles := copyFile(selectedFiles) // Copy files from note directory to current directory + minor file editing
	_ = copyImages(copiedFiles)            // Copy all images in the copied files to current directory

	message := mdToHTML(copiedFiles) // Convert copied markdown files to HTML

	contentLock.Lock()
	webContent = string(message) // Write HTML file to webContent
	//sendEmail()                  // Send notification email with 'Time to Review' message
	contentLock.Unlock()
}

// Fills and returns slice containing all files in directory specified
func repopulate(dir string) []string {
	res := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(fmt.Errorf("Error occured: %v\n", err))
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

// Selects a random file from each of the n directories specified by user in the command line
func smartSelect() []string {
	var result []string // store n random filepaths in here (n = # specified in command line args)
	var totalSize int64 = 0

	for dirName, dirList := range dirMap {
		if len(dirList) == 0 {
			// This should never happen because smartSelect is always called after potential repopulates
			fmt.Printf("Failed on iteration for %s\n", dirName)
			os.Exit(1)
		}

		randNum := rand.Intn(len(dirList)) // Generate random number between 0 and num of files
		chosen := dirList[randNum]

		info, err := os.Stat(chosen)
		if err != nil {
			fmt.Printf("Problem obtaining the file info of the file: %s\n", chosen)
			continue
		}

		totalSize += info.Size()

		// If the totalSize of the files we wish to add is larger than 20 kB, don't add the additional file
		if len(result) > 1 && totalSize > 20480 {
			totalSize -= info.Size()
			break
		}

		for i, val := range dirList {
			if val == chosen {
				dirList = append(dirList[:i], dirList[i+1:]...)
				dirMap[dirName] = dirList
				break
			}
		}

		result = append(result, chosen)
	}

	//fmt.Printf("Chosen files: %s. Total size: %d.\n", result, totalSize)
	return result
}

// Sends the notification email with the information specified in .env file
func sendEmail() {
	err := godotenv.Load(".env")

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
	m.SetHeader("Subject", "Dio: It's Time to Extend Your Knowledge")

	// Get IP address of the server
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("Unable to ressolve server's IP address")
	}
	defer conn.Close()
	addrWPort := conn.LocalAddr().String()
	splitted := strings.Split(addrWPort, ":")

	message := fmt.Sprintf(`<h1>The Road to Excellence is Paved in the Mundance</h1>
		<p>You are on the right path: %s%s</p>`, splitted[0], PORT)
	m.SetBody("text/html", message)

	d := gomail.NewDialer(smtpHost, 587, from, mailPass)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("Email sent successfully!")
}

// Deletes the output directory and everything in it, then recreates it and the images directory
func cleanRepo() {
	err := os.RemoveAll("output")
	if err != nil {
		fmt.Println("Unable to clean the current working directory.")
		os.Exit(69)
	}

	err = os.MkdirAll("output/images", 0777)
	if err != nil {
		fmt.Println("Unable to clean the current working directory.")
		os.Exit(420)
	}
}

// Read all specifed MD files in files and render as one entire HTML file
func mdToHTML(files []string) []byte {
	// read and store markdown file in 'mdContent'
	var contents []byte
	// Read each file and concatenate them all together to form a large file
	for i := range files {
		readData, err := os.ReadFile("output/" + files[i])
		if err != nil {
			fmt.Println(err)
		} else {
			contents = append(contents, readData...)
		}
	}

	// Create a parser and parse the md stored as bytes
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(contents)

	// Create a html renderer
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	// Take the parsed MD and turn it to HTML, surrounding it with the below boilerplate,
	// loading the Mathjax parser for Latex equations, and 'style.css' for mild styling
	res := fmt.Sprintf(`<!DOCTYPE html>
	<html lang="en">

	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>Dio: Greatness in the Mundane</title>
		<link rel="stylesheet" href="/style.css">
	</head>

	<body>
	<script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
		<div class="content">
		%s
		</div>
	</body>

	</html>`, markdown.Render(doc, renderer))

	// Write content to final.html for debug / visibility purposes
	os.WriteFile("final.html", []byte(res), 0777)
	return []byte(res)
}

// Read contents of a file specified in selectedFiles, and copy it into cwd
func copyFile(selectedFiles []string) []string {
	var copiedFiles []string

	// Copy the files pre-selected files
	for _, fullPathName := range selectedFiles {
		splitted := strings.Split(fullPathName, "/")
		fileName := splitted[len(splitted)-1]

		// Open original markdown file
		src, err := os.Open(fullPathName)
		if err != nil {
			fmt.Println(fmt.Errorf("Error occured: %v\n", err))
			os.Exit(1)
		}
		defer src.Close()
		// Create file that will hold the original's contents
		dst, err := os.Create("output/" + fileName)
		if err != nil {
			fmt.Println(fmt.Errorf("Error occured: %v\n", err))
			os.Exit(2)
		}
		defer dst.Close()

		// Copy contents of the original into the newly created file
		_, err = io.Copy(dst, src)
		if err != nil {
			fmt.Println(fmt.Errorf("Error occured: %v\n", err))
			os.Exit(3)
		}

		// Add newlines to the end of every copied file (to easily visually distinguish the files)
		_, err = dst.WriteString("\n\n\n\n")
		if err != nil {
			fmt.Printf("Failed to write newline to the file: '%s'.\n", fileName)
			os.Exit(4)
		}

		// Keep a record of the files that were copied
		copiedFiles = append(copiedFiles, fileName)
		fmt.Println("Successfully copied: ", fileName)
	}

	//fmt.Println("Copied files: ", copiedFiles)
	return copiedFiles
}

// Go through list of files, and copy all images in each file
func copyImages(selectedFiles []string) []string {
	var copiedImages []string

	for _, fileName := range selectedFiles {
		extracted := extractPhotos(fileName)
		copiedImages = append(copiedImages, extracted...)
		//fmt.Println("Extracted: ", extracted)

		for _, photoName := range extracted {
			// Open image file for reading
			src, err := os.Open(imgDir + "/" + photoName)
			if err != nil {
				fmt.Println(fmt.Errorf("Error occured: %v\n", err))
			}
			defer src.Close()
			// Create image file that will hold the copy of the original image's contents
			dst, err := os.Create("output/images/" + photoName)
			if err != nil {
				fmt.Println(fmt.Errorf("Error occured: %v\n", err))
				os.Exit(5)
			}
			defer dst.Close()

			// Copy contents of the original into the newly created file
			_, err = io.Copy(dst, src)

			//if err != nil {
			//	fmt.Printf("Failed to copy file '%s' contents: %v\n", photoName, err)
			//} else {
			//	fmt.Printf("Copy of %s created successfully!\n", photoName)
			//}
		}
	}

	//fmt.Printf("Copied Images: %s\n\n", copiedImages)
	return copiedImages
}

// Use regex to extract names of images from a file and replace them with the new location of the image file
func extractPhotos(fileName string) []string {
	var allPhotos []string
	var content []byte
	var err error

	if strings.Split(fileName, "/")[0] == "test_input" {
		content, err = os.ReadFile(fileName)
	} else {
		content, err = os.ReadFile("output/" + fileName)
	}
	if err != nil {
		fmt.Println(fmt.Errorf("Error occured: %v\n", err))
		os.Exit(7)
	}

	re := regexp.MustCompile(`!\[\[(.+?\|?.*?)\]\]|!\[\]\((.+?)\)`)

	result := re.ReplaceAllStringFunc(string(content), func(match string) string {
		matches := re.FindStringSubmatch(match)
		var s string

		// Check the first capture group (for ![[...]])
		if len(matches) > 1 && matches[1] != "" {
			// If there is a '|', split the string and take the first part
			parts := regexp.MustCompile(`^[^|]+`).FindStringSubmatch(matches[1])
			if len(parts) > 0 {
				s = parts[0]
			}
		}
		// Otherwise, check the second capture group (for ![](...))
		if len(matches) > 2 && matches[2] != "" {
			s = matches[2]
		}

		s = regexp.MustCompile(`%20`).ReplaceAllString(s, " ")

		allPhotos = append(allPhotos, s)

		return fmt.Sprintf("![](output/images/%s)", s)
	})

	err = os.WriteFile("output/"+fileName, []byte(result), 0644)
	if err != nil {
		fmt.Printf("Error writing to file %s\n", fileName)
		os.Exit(9)
	}

	return allPhotos
}
