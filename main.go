package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const cachefolder = "buildprog"
const buildFile = "build.txt"

var buildArgs = []string{"build", "-work", "-a", "-p", "1", "-x"}

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

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("buildprog [-cleancache] Go compiler arguments")
	}

	start := 1
	if os.Args[1] == "-cleancache" {
		err = os.RemoveAll(filepath.Join(cachedir, cachefolder))
		handle(err)
		start++
	}
	flags := os.Args[start:]

	path := filepath.Join(cachedir, cachefolder, currdir)
	err = os.MkdirAll(path, os.ModePerm)
	handle(err)

	outFile := filepath.Join(path, buildFile)
	cmd := exec.Command("go", append(buildArgs, flags...)...)

	errout := bytes.NewBufferString("")
	cmd.Stderr = errout

	if _, err := os.Stat(outFile); os.IsNotExist(err) {
		fmt.Println("No analysis graph available, compiling without progress this time...")
		out, err := os.Create(outFile)
		handle(err)
		cmd.Stdout = out
	} else {
		fmt.Println("unimplemented!")
	}
	err = cmd.Run()
	if err != nil {
		handle(errors.New(errout.String()))
	}
}
