package main

import (
	"os/exec"

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
	exec.Command("make", "-C", "uuid-sample", "clean").Run()
	exec.Command("make", "-C", "mod_wasm", "clean").Run()
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

	r.Step(demo.S(
		"Build uuid-sample",
	), demo.S(
		"make -C uuid-sample",
	))

	r.Step(demo.S(
		"uuid-sample artifacts",
	), demo.S(
		"tree uuid-sample",
	))

	r.Step(demo.S(
		"Run uuid.wasm",
	), demo.S(
		"wasmtime uuid-sample/uuid.wasm",
	))

	r.Step(demo.S(
		"Run uuid-cgi.wasm",
	), demo.S(
		"wasmtime uuid-sample/uuid-cgi.wasm",
	))

	r.Step(demo.S(
		"Run uuid-wws.wasm",
	), demo.S(
		"wasmtime uuid-sample/uuid-wws.wasm",
	))

	r.Step(demo.S(
		"Run uuid with wws",
	), demo.S(
		"wws uuid-sample",
	))

	r.Step(demo.S(
		"Run uuid with spin",
	), demo.S(
		"bat spin/spin.toml",
	))

	r.Step(demo.S(
		"Run uuid with spin",
	), demo.S(
		"spin up -f spin",
	))

	r.Step(demo.S(
		"Run uuid in Apache with mod_wasm",
	), demo.S(
		"grep -A6 'IfModule wasm_module' mod_wasm/conf/httpd.conf",
	))

	r.Step(demo.S(
		"Run uuid in Apache with mod_wasm",
	), demo.S(
		"make -C mod_wasm run",
	))

	// TODO: Add PHP button generator demo

	return r
}
