package commands

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var OBJECT_DIR string = ".ggit/objects"

func HashObject(args []string) {
	if len(args) == 0 {
		fmt.Println("hash-object command requires a file name")
		os.Exit(1)
	}

	fileContent := open(args[0])
	oid := hash(fileContent)
	objectPath := filepath.Join(OBJECT_DIR, oid)
	write(objectPath, fileContent)
	fmt.Println(oid)
}

func open(fileName string) []byte {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, fileName)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Unable to open %s\n", filePath)
	}
	return file
}

func hash(fileContent []byte) string {
	sum := sha256.Sum256(fileContent)
	return fmt.Sprintf("%x", sum)

}

func write(objectPath string, fileContent []byte) {
	os.WriteFile(objectPath, fileContent, 0644)
}
