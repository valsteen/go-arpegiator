package runner

import (
	"bufio"
	"go-arpegiator/services"
	"os"
)

func Sleep() {
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	services.MustNot(err)
}
