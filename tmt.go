package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/schollz/progressbar/v2"
)

var (
	m = flag.Int("m", 25, "")
	c = 0
)

var usage = `Usage: tmt [options...]

Options:
	-m	Number of minutes to count. Default is 25 minutes.
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}

	flag.Parse()

	end := *m * 60
	bar := progressbar.New(end)

	bar.RenderBlank()
	for range time.Tick(1 * time.Second) {
		if c > end {
			break
		}
		bar.Add(1)
		c++
	}

}
