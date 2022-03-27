package commands

import (
	"bytes"
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
	object := settype(fileContent)
	oid := hash(fileContent)
	objectPath := filepath.Join(OBJECT_DIR, oid)
	write(objectPath, object)
	return fmt.Sprintf("%s\n", oid), nil
}

func CatFile(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("cat-file required the object id")
	}
	objectPath := filepath.Join(OBJECT_DIR, args[0])
	fileContent := openObject(objectPath)
	return fmt.Sprint(string(fileContent[:])), nil
}

func settype(content []byte) []byte {
	defaultType := []byte("blob")
	nullByte := byte(0)
	header := append(defaultType, nullByte)
	return append(header, content...)
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

func openObject(fileName string) []byte {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, fileName)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Unable to open %s\n", filePath)
	}
	_, after, _ := bytes.Cut(file, []byte{0})
	return after
}

func hash(fileContent []byte) string {
	sum := sha256.Sum256(fileContent)
	return fmt.Sprintf("%x", sum)
}

func write(objectPath string, fileContent []byte) {
	os.WriteFile(objectPath, fileContent, 0644)
}
