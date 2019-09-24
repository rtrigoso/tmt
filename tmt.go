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
	r = flag.Int("r", 5, "")
	n = flag.Int("n", 1, "")
	c = 0
	p = [2]string{"WORK", "REST"}
)

var usage = `Usage: tmt [options...]

Options:
	-m	Work length in minutes. Defaults to 25 minutes.
	-r	Rest length in minutes. Defaults to 5 minutes.
	-n	Number of sets to run for. Defaults to 1.
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}
	flag.Parse()

	for j := 0; j < *n; j++ {
		for _, label := range p {
			switch label {
			case "WORK":
				startProgress(*m, label)
				break
			case "REST":
				startProgress(*r, label)
				break
			default:
				os.Exit(1)
			}
		}
	}

	roundsMessage := "rounds"
	if *n == 1 {
		roundsMessage = "round"
	}

	workMessage := "minutes"
	if *m == 1 {
		workMessage = "minute"
	}

	restMessage := "minutes"
	if *r == 1 {
		restMessage = "minute"
	}

	fmt.Printf("%d %s of %d %s work and %d %s rest pomodoro finished\n", *n, roundsMessage, *m, workMessage, *r, restMessage)
}

func startProgress(t int, l string) {
	t *= 60
	bar := progressbar.NewOptions(t, progressbar.OptionSetDescription(l))
	bar.RenderBlank()

	var c int
	for range time.Tick(1 * time.Second) {
		if c > t {
			bar.Clear()
			break
		}
		bar.Add(1)
		c++
	}
}
