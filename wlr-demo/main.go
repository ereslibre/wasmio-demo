package main

import (
	demo "github.com/saschagrunert/demo"
	"github.com/urfave/cli/v2"
)

func main() {
	d := demo.New()

	d.Setup(setup)
	d.Cleanup(cleanup)

	d.Add(runUUID(), "run-uuid", "Run uuid")

	d.Run()
}

func setup(ctx *cli.Context) error {
	return cleanup(ctx)
}

func cleanup(ctx *cli.Context) error {
	return nil
}

func runUUID() *demo.Run {
	r := demo.NewRun(
		"Run UUID",
	)

	r.Step(demo.S(
		"Install library bundle",
	), demo.S(
		"tree libs",
	))

	r.Step(demo.S(
		"Set up PKG_CONFIG_PATH",
	), demo.S(
		"tree $PKG_CONFIG_PATH",
	))

	r.Step(demo.S(
		"Set up PKG_CONFIG_SYSROOT_DIR",
	), demo.S(
		"echo $PKG_CONFIG_SYSROOT_DIR",
	))

	r.Step(demo.S(
		"Check that pkg-config detects configuration files",
	), demo.S(
		"pkg-config --list-all",
	))

	return r
}
