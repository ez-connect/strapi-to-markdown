package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const appName = "strapi-to-markdown"
const buildDir = "build"

var platforms = []string{"linux", "darwin", "windows"}

const arch = "amd64"

func main() {
	for _, v := range platforms {
		fmt.Println(appName, ":", v, arch)
		os.Setenv("GOOS", v)
		os.Setenv("GOARCH", arch)

		output := fmt.Sprintf("%s/%s-%s-%s", buildDir, appName, v, arch)
		cmd := exec.Command("go", "build", "-o", output, "-ldflags", "-s -w")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
