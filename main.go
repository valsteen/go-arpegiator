package main

import (
	"go-arpegiator/runner"
)

func main() {
	runArpInDevice := runner.RunArpInDevice("Arp")
	defer runArpInDevice.Close()
	runner.Sleep()
}
