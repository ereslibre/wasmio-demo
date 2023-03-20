package main

import (
	"os"

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
	return cleanup(ctx)
}

func cleanup(ctx *cli.Context) error {
	os.Remove("wws")
	return nil
}

func installAndRunWWS() *demo.Run {
	r := demo.NewRun(
		"Install and run Wasm Workers Server (wws)",
	)

	r.Step(demo.S(
		"Install Wasm Workers Server (wws)",
	), demo.S(
		"curl -fsSL https://workers.wasmlabs.dev/install |",
		"bash -s -- --local",
	))

	r.Step(demo.S(
		"Check wws help",
	), demo.S(
		"./wws --help",
	))

	r.Step(demo.S(
		"Run wws",
	), demo.S(
		"./wws .",
	))

	r.Step(demo.S("Check documentation at https://workers.wasmlabs.dev/"), nil)

	return r
}
