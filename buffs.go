package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/schollz/progressbar/v3"
)

type lenBuff struct {
	len int
}

func (l *lenBuff) Write(p []byte) (int, error) {
	l.len += len(p)
	return len(p), nil
}

type progBuff struct {
	origLen  int
	len      int
	finished bool
	bar      *progressbar.ProgressBar
}

func (p *progBuff) Write(d []byte) (int, error) {
	if !p.finished {
		p.len += len(d)
		if p.len > p.origLen {
			handle(p.bar.Finish())
			p.finished = true
			fmt.Println()
			fmt.Println("It appears a dependency was added. Progress is no longer available, the cache will be updated.")
		}
		p.bar.Add(len(d))
	}
	return len(d), nil
}

func (p *progBuff) Close(outFile string) {
	if p.finished {
		err := os.WriteFile(outFile, []byte(strconv.Itoa(p.len)), os.ModePerm)
		handle(err)
	} else {
		handle(p.bar.Finish())
		p.finished = true
		fmt.Println()
	}
}
