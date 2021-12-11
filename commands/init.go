package commands

import (
	"fmt"
	"os"
)

var GIT_DIR string = ".ggit"

func Init(args []string) {
	wd, _ := os.Getwd()
	os.Mkdir(fmt.Sprintf("%s/%s", wd, GIT_DIR), os.ModeDir)
	fmt.Printf("Initialized empty ggit repository in %s/%s\n", wd, GIT_DIR)
}
