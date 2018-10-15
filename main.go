package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/wouterbeets/cyclist"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		os.Exit(1)
	}
	c, err := cyclist.New(map[string]cyclist.Cycle{
		"0": cyclist.Cycle{
			Entries: []cyclist.Entry{
				{Cmd: "xrandr --auto"},
			},
		},
	})
	if err != nil {
		os.Exit(1)
	}
	nbScreens := args[0]
	if err != nil {
		os.Exit(1)
	}
	cmd, err := c.Cycle(nbScreens)
	comm := exec.Command("notify-send", cmd)
	err = comm.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
