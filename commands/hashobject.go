package commands

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var OBJECT_DIR string = ".ggit/objects"

func HashObject(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("hash-object command requires a file name")
	}

	fileContent := open(args[0])
	oid := hash(fileContent)
	objectPath := filepath.Join(OBJECT_DIR, oid)
	write(objectPath, fileContent)
	return fmt.Sprintf("%s\n", oid), nil
}

func CatFile(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("cat-file required the object id")
	}
	objectPath := filepath.Join(OBJECT_DIR, args[0])
	fileContent := open(objectPath)
	return fmt.Sprint(string(fileContent[:])), nil
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
