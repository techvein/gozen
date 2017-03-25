package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	os.Chdir(filepath.Join(os.Getenv("GOPATH"), "src/gozen"))
	_, err := exec.LookPath("dep")
	if err != nil {
		log.Fatal("Please install dep that is a package management tool.")
	}
	log.Println("done")

	cmd := exec.Command("dep", "ensure", "-update", "-v")
	fmt.Print("run", cmd.Path, " ensure -update -v ...")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done dep ensure update.")

}
