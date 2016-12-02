package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

var gopath string

var libs = []string{
	"github.com/kardianos/govendor",
	"bitbucket.org/liamstask/goose/cmd/goose",
	"gopkg.in/yaml.v2",
	"github.com/google/logger",
}

func main() {
	gopath = os.Getenv("GOPATH")
	os.Chdir(filepath.Join(gopath, "src/gozen"))

	goget(libs)
}

func goget(libs []string) {
	var wg sync.WaitGroup
	for _, lib := range libs {
		wg.Add(1)
		go func(lib string) {
			defer wg.Done()
			cmd := exec.Command("go", "get", lib)
			fmt.Println("Run", cmd.Args)
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Done", cmd.Args)
		}(lib)
	}
	wg.Wait()
}
