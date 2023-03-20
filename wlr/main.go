package main

import (
	demo "github.com/saschagrunert/demo"
	"github.com/urfave/cli/v2"
)

func main() {
	d := demo.New()

	d.Setup(setup)
	d.Cleanup(cleanup)

	d.Add(runPHP(), "run-php", "Run PHP")

	d.Run()
}

func setup(ctx *cli.Context) error {
	return cleanup(ctx)
}

func cleanup(ctx *cli.Context) error {
	return nil
}

func runPHP() *demo.Run {
	r := demo.NewRun(
		"Run PHP",
	)

	return r
}
