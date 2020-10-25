package main

import (
	"go-arpegiator/runner"
)

func main() {
	//deviceRunner := runner.RunDevice("Arp", func(notes devices.Notes) { fmt.Println(notes) })
	//defer deviceRunner.Close()

	runArpInDevice := runner.RunArpInDevice("Arp")
	defer runArpInDevice.Close()
	runner.Sleep()
}
