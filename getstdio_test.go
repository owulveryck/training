package training_test

import (
	"log"
	"os"
	"os/exec"

	"github.com/owulveryck/training"
)

func ExampleGetStdIO() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	training.ListenAddr = "localhost:1234"
	stdin, stdout, stderr, err := training.GetStdIO(training.Network)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Stdin = <-stdin
	cmd.Stdout = <-stdout
	cmd.Stderr = <-stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
