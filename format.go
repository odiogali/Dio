package main

import (
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"fmt"
)

// Read all specifed MD files in files and render as one entire HTML file
func mdToHTML(files []string) []byte {
	// read and store markdown file in 'mdContent'
	var contents []byte
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

	// Create a html renderer and from the parsed markdown, render html
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	os.WriteFile("final.html", markdown.Render(doc, renderer), 0777)

	return markdown.Render(doc, renderer)
}

// Read contents of a file specified in selectedFiles, and copy it into cwd
func copyFile(selectedFiles []string) []string {
	var copiedFiles []string

	// Copy the files selected by smartSelect
	for _, fullPathName := range selectedFiles {
		splitted := strings.Split(fullPathName, "/")
		fileName := splitted[len(splitted)-1]

		// Open original markdown file
		src, err := os.Open(fullPathName)
		if err != nil {
			fmt.Printf("Error opening the file: %s\n", fullPathName)
			os.Exit(1)

		}
		defer src.Close()
		// Create file the copy of the original will be stored in
		dst, err := os.Create("output/" + fileName)
		if err != nil {
			fmt.Printf("Failed to create destination file: %s due to %v\n", fileName, err)
			os.Exit(2)
		}
		defer dst.Close()

		// Copy contents of the original into the newly created file
		_, err = io.Copy(dst, src)
		if err != nil {
			fmt.Printf("Failed to copy file '%s' contents: %v\n", fileName, err)
			os.Exit(3)
		}

		_, err = dst.WriteString("\n")
		if err != nil {
			fmt.Printf("Failed to write newline to the file: '%s'.\n", fileName)
			os.Exit(4)
		}

		copiedFiles = append(copiedFiles, fileName)
		fmt.Printf("Copy of %s created successfully!\n", fileName)
	}

	return copiedFiles
}

// Read contents of a image file specified in selectedFiles, and copy it into cwd
func copyImages(selectedFiles []string) []string {
	var copiedImages []string

	for _, fileName := range selectedFiles {
		extracted := extractPhotos(fileName)
		copiedImages = append(copiedImages, extracted...)

		fmt.Printf("Extracted: ")
		fmt.Println(extracted)
		for _, photoName := range extracted {
			// Open image for reading
			src, err := os.Open(ImgDir + "/" + photoName)
			if err != nil {
				fmt.Printf("Error opening the file: %s, due to error: %v\n", photoName, err)
			}
			defer src.Close()
			// Create file the copy of the original will be stored in
			dst, err := os.Create("output/images/" + photoName)
			if err != nil {
				fmt.Printf("Failed to create image file: %s due to %v\n", photoName, err)
				os.Exit(5)
			}
			defer dst.Close()

			// Copy contents of the original into the newly created file
			_, err = io.Copy(dst, src)
			if err != nil {
				fmt.Printf("Failed to copy file '%s' contents: %v\n", photoName, err)
			} else {
				fmt.Printf("Copy of %s created successfully!\n", photoName)
			}
		}
	}

	fmt.Printf("Copied Images: %s\n\n", copiedImages)

	return copiedImages
}

// Use regex to extract names of images from a file and replace them with the new location of the image file
func extractPhotos(fileName string) []string {
	var allPhotos []string

	content, err := os.ReadFile("output/" + fileName)
	if err != nil {
		fmt.Printf("Error reading file '%s', due to %v\n", fileName, err)
		os.Exit(7)
	}

	re := regexp.MustCompile(`!\[\[(.+?)\]\]|!\[\]\((.+?)\)`)

	result := re.ReplaceAllStringFunc(string(content), func(match string) string {
		matches := re.FindStringSubmatch(match)
		var s string

		if len(matches) > 1 && matches[1] != "" {
			s = matches[1]
		} else if len(matches) > 2 && matches[2] != "" {
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
