package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var skip struct {
	files []string
	dirs  []string
}

func init() {
	skip.files = []string{
		"bundle.js",
		"DS_Store",
	}

	skip.dirs = []string{
		"node_modules",
		".git",
		".idea",
		".vscode",
		".svn",
	}
}

func main() {
	rename := flag.Bool("file", false, "Search and replace file name")
	text := flag.Bool("text", false, "Search and replace text")
	search := flag.String("search", "", "Text to search for")
	replace := flag.String("replace", "", "Text to replace with")
	dir := flag.String("dir", "", "Directory to search")
	flag.Parse()

	if !*rename && !*text {
		flag.PrintDefaults()
		log.Fatal("Must specify either -file or -text")
	}

	if *rename {
		err := SearchAndRenameFile(*dir, *search, *replace)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *text {
		err := SearchAndReplaceText(*dir, *search, *replace)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Done")
}

func SearchAndReplaceText(rootDir string, textToFind string, textToReplace string) error {
	// Recursively search a directory and replace text in all found files.

	var files []string

	err := filepath.Walk(rootDir, visit(&files))
	if err != nil {
		return err
	}

	for _, filepath := range files {
		err := ReplaceTextInFile(filepath, textToFind, textToReplace)
		if err != nil {
			return err
		}
	}

	return nil
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() && isSkipDir(info.Name()) {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			if !isSkipFile(info.Name()) {
				*files = append(*files, path)
			}
		}
		return nil
	}
}

func ReplaceTextInFile(filepath string, textToFind string, textToReplace string) error {
	// Replaces all occorances of text with new text in file.

	file, _ := ioutil.ReadFile(filepath)
	rf := []byte(strings.ReplaceAll(string(file), textToFind, textToReplace))
	err := ioutil.WriteFile(filepath, rf, 0666)

	return err
}

func SearchAndRenameFile(rootDir string, oldName string, newName string) error {
	// Recursively search directory for files named {oldName} and rename to {newName}

	var files []string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() && isSkipDir(info.Name()) {
			return filepath.SkipDir
		}
		if info.IsDir() && !isSkipDir(info.Name()) {
			for _, path := range files {
				dir, fileName := filepath.Split(path)
				if fileName == oldName {
					err = os.Rename(path, dir+newName)
				}
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func isSkipDir(dir string) bool {
	for _, skipDir := range skip.dirs {
		if dir == skipDir {
			return true
		}
	}
	return false
}

func isSkipFile(file string) bool {
	for _, skipFile := range skip.files {
		if file == skipFile {
			return true
		}
	}
	return false
}
