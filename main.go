package main

import (
	"fmt"
	"go-arpegiator/devices"
	"go-arpegiator/runner"
)

func main() {
	deviceRunner := runner.RunDevice("Arp", func(notes devices.Notes) { fmt.Println(notes) })
	defer deviceRunner.Close()
	runner.Sleep()
}
