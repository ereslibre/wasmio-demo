package main

import (
	"os"
	"os/exec"

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
	exec.Command("pkill", "-x", "npm").Run()
	return cleanup(ctx)
}

func cleanup(ctx *cli.Context) error {
	exec.Command("pkill", "-x", "wws").Run()
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
		"Workers subtree",
	), demo.S(
		"tree my-workers",
	))

	r.Step(demo.S(
		"Javascript worker code",
	), demo.S(
		"bat my-workers/js-basic/index.js",
	))

	r.Step(demo.S(
		"Run wws",
	), demo.S(
		"./wws my-workers &",
	))

	r.Step(demo.S("Check documentation at https://workers.wasmlabs.dev/"), nil)

	r.Step(demo.S("Let's run the same worker code, but in Cloudflare Workers"), nil)

	r.Step(demo.S(
		"Cloudflare Javascript worker code",
	), demo.S(
		"bat cf-workers/src/index.js",
	))

	r.Step(demo.S(
		"Run wrangler",
	), demo.S(
		"cd cf-workers &&",
		"npx wrangler dev -l &",
	))

	return r
}
