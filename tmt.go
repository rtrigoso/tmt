package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/schollz/progressbar/v2"
)

var (
	m   = flag.Int("m", 25, "")
	r   = flag.Int("r", 5, "")
	n   = flag.Int("n", 1, "")
	x   = flag.String("x", "", "")
	c   = 0
	p   = [2]string{"WORK", "REST"}
	err error
)

var usage = `Usage: tmt [options...]

Options:
	-m	Work length in minutes. Defaults to 25 minutes.
	-r	Rest length in minutes. Defaults to 5 minutes.
	-n	Number of sets to run for. Defaults to 1.
	-x	Quoted command to run after the last set is done, instead of an end message
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}
	flag.Parse()

	defer func() {
		if err != nil {
			fmt.Printf("error running tmt: %v", err)
			os.Exit(1)
		}
	}()

	for j := 0; j < *n; j++ {
		for _, label := range p {
			switch label {
			case "WORK":
				startProgress(*m, "[green]"+label)
				break
			case "REST":
				startProgress(*r, "[red]"+label)
				break
			default:
				os.Exit(1)
			}
		}
	}

	if *x != "" {
		cmd := exec.Cmd{}
		cmd, err = cmdBuild(*x)

		var b []byte
		b, err = cmd.CombinedOutput()
		fmt.Printf(string(b))
		os.Exit(0)
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
	bar := progressbar.NewOptions(t, progressbar.OptionSetDescription(l), progressbar.OptionEnableColorCodes(true))
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

func cmdBuild(s string) (exec.Cmd, error) {
	a := strings.Fields(s)
	name := a[0]

	args := []string{name}
	if len(a) > 1 {
		for _, f := range a[1:] {
			args = append(args, f)
		}
	}

	cmd := exec.Cmd{
		Path: name,
		Args: args,
	}

	if filepath.Base(name) == name {
		if lp, err := exec.LookPath(name); err != nil {
			return cmd, err
		} else {
			cmd.Path = lp
		}
	}

	return cmd, nil
}
