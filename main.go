package main

import (
	"go-arpegiator/runner"
)

func main() {
	devices := runner.RunArpegiator("NotesIn", "Arp")
	defer devices.Close()
	runner.Sleep()
}
