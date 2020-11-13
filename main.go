package main

import (
	"go-arpegiator/runner"
)

func main() {
	devices := runner.RunArpegiator("NotesIn", "Pattern")
	defer devices.Close()
	runner.Sleep()
}
