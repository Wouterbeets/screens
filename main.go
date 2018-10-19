package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/wouterbeets/cyclist"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("error: no id")
		os.Exit(1)
	}
	c, err := cyclist.New(
		map[string]cyclist.Cycle{
			"0": cyclist.Cycle{
				Entries: []string{
					"xrandr --auto",
				},
			},
		},
		os.Getenv("HOME")+"/.screens",
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	nbScreens := args[0]
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd, err := c.Cycle(nbScreens)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	comm := exec.Command("notify-send", cmd)
	err = comm.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmdArgs := strings.Split(cmd, " ")

	comm = exec.Command(cmdArgs[0], cmdArgs[1:]...)
	err = comm.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
