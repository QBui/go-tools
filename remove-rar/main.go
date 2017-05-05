package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func visit(path string, f os.FileInfo, err error) error {
	//if f.IsDir() && containsFileExtension(path, ".rar") && containsFileExtension(path, ".iso") {
	if f.IsDir() && containsFileExtension(path, ".rar") && containsFileExtension(path, ".mp4") {
		fmt.Printf("Found: %s\n", path)
		deleteRARFiles(path)
	}
	return nil
}

func containsFileExtension(path string, extension string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == extension {
				return true
			}
		}
	}
	return false
}

func deleteRARFiles(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`\.r..`)
	for _, file := range files {
		if file.Mode().IsRegular() {
			match := re.FindString(filepath.Ext(file.Name()))
			if len(match) > 0 {
				//os.Remove(filepath.Join(path, file.Name()))
				fmt.Printf("Deleted: %#v %v\n", match, file.Name())
			} else {
				fmt.Printf("Not matched: %#v %v\n", match, file.Name())
			}
		}
	}
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}
