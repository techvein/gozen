package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var tools = []string{
	"goget",
	"setup",
}

func main() {
	gopath := os.Getenv("GOPATH")
	err := os.Chdir(filepath.Join(gopath, "src/tools"))
	if err != nil {
		log.Println(err)
	}

	goRunTools()

}

// toolsを実行する。
func goRunTools() {
	for _, tool := range tools {
		toolPath := filepath.Join("tools", tool+".go")
		cmd := exec.Command("go", "run", toolPath)
		fmt.Print("go run ", toolPath, "...")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Done.")
	}
}
