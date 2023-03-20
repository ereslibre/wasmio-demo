.PHONY: wws
wws:
	@clear
	@nix develop ~/projects/nixities#temporary.wasmio.wws --command go run . --install-and-run-wws

.PHONY: shell
shell:
	nix develop ~/projects/nixities#temporary.wasmio.wws

.PHONY: clean
clean:
	rm wws
