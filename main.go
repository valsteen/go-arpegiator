package main

import (
	"go-arpegiator/runner"
)

func main() {
	runArpInDevice := runner.RunArpegiator("NotesIn", "Arp")
	defer runArpInDevice.Close()
	runner.Sleep()
}
