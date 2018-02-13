package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var inputFolderPath, inputFilePath string
	var mobijakeList []string
	flag.StringVar(&inputFolderPath, "i", "", "Path to folder with files")
	flag.StringVar(&inputFilePath, "a", "", "Path to file with archive downloads")

	flag.Parse()

	if inputFolderPath == "" {
		fmt.Println("Input folder with mobijake files has to be specified i.e. with -i path/to/folder")
		return
	}
	if inputFilePath == "" {
		fmt.Println("Input file with archive download has to be specified i.e. with -a path/to/file")
		return
	}

	inFile, _ := os.Open(inputFilePath)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	files, err := ioutil.ReadDir(inputFolderPath)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		mobijake := strings.Split(scanner.Text(), " ")[1]
		mobijakeList = append(mobijakeList, mobijake)
	}

	for _, file := range files {
		fileName := file.Name()
		replace(mobijakeList, fileName, inputFolderPath)
	}
}

func replace(mobijakes []string, fileName string, folder string) {
	for _, mj := range mobijakes {
		if strings.Contains(fileName, mj) {
			newName := strings.Replace(fileName, "-"+mj, "", -1)
			originalPath := folder + "/" + fileName
			newPath := folder + "/" + newName
			err := os.Rename(originalPath, newPath)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
