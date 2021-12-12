package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

var GIT_DIR string = ".ggit"

func Init(args []string) {
	wd, _ := os.Getwd()
	os.Mkdir(filepath.Join(wd, GIT_DIR), 0755)
	os.Mkdir(filepath.Join(wd, GIT_DIR, "objects"), 0755)

	fmt.Printf("Initialized empty ggit repository in %s/%s\n", wd, GIT_DIR)
}
