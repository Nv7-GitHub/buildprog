package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/schollz/progressbar/v3"
)

const cachefolder = "buildprog"
const buildFile = "build.txt"

var buildArgs = []string{"build", "-work", "-p", "1", "-x"}

func handle(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func main() {
	currdir, err := os.Getwd()
	handle(err)
	cachedir, err := os.UserCacheDir()
	handle(err)

	start := 1
	// Flags
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Println("buildprog [-cleancache] Go compiler arguments")
			return
		}

		if os.Args[1] == "-cleancache" {
			fmt.Println("Removing cache...")
			err = os.RemoveAll(filepath.Join(cachedir, cachefolder))
			handle(err)
			start++
		}
	}

	flags := make([]string, 0)
	if len(os.Args) > start {
		flags = os.Args[start:]
	}

	path := filepath.Join(cachedir, cachefolder, currdir)
	err = os.MkdirAll(path, os.ModePerm)
	handle(err)

	outFile := filepath.Join(path, buildFile)
	cmd := exec.Command("go", append(buildArgs, flags...)...)

	if _, err := os.Stat(outFile); os.IsNotExist(err) {
		fmt.Println("No analysis graph available, compiling without progress this time...")

		buff := &lenBuff{}
		cmd.Stderr = buff
		err = cmd.Run()
		handle(err)

		err = os.WriteFile(outFile, []byte(strconv.Itoa(buff.len)), os.ModePerm)
		handle(err)
	} else {
		fmt.Println("Loading cache...")
		cache, err := os.ReadFile(outFile)
		handle(err)
		prevlen, err := strconv.Atoi(string(cache))
		handle(err)

		bar := progressbar.New(prevlen)
		buff := &progBuff{
			origLen: prevlen,
			bar:     bar,
			len:     0,
		}

		fmt.Println("Building...")

		cmd.Stderr = buff
		err = cmd.Run()
		handle(err)

		buff.Close(outFile)
	}
	fmt.Println("Successfully built!")
}
