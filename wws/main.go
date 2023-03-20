package main

import (
	demo "github.com/saschagrunert/demo"
	"github.com/urfave/cli/v2"
)

func main() {
	d := demo.New()

	d.Setup(setup)
	d.Cleanup(cleanup)

	d.Add(installAndRunWWS(), "install-and-run-wws", "Installs and runs Wasm Workers Server")

	d.Run()
}

func setup(ctx *cli.Context) error {
	return nil
}

func cleanup(ctx *cli.Context) error {
	return nil
}

func installAndRunWWS() *demo.Run {
	r := demo.NewRun(
		"Install and run Wasm Workers Server (wws)",
	)

	r.Step(demo.S(
		"Install Wasm Workers Server (wws)",
	), demo.S(
		"curl -fsSL https://workers.wasmlabs.dev/install | bash -s -- --local",
	))

	return r
}
