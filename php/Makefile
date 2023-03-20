image.png:
	wasmtime run --allow-unknown-exports --dir . php-cgi-8.2.0.wasm -- index.php 2> /dev/null | tail -n+6 > $@

.PHONY: clean
clean:
	rm *.png
