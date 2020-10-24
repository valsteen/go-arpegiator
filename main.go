package main

import (
	"go-arpegiator/runner"
)

func main() {
	deviceRunner := runner.RunDevice()
	defer deviceRunner.Close()
	runner.Sleep()
}
