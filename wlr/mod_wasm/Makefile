.PHONY: build
build:
	docker build -t mod_wasm:latest .

.PHONY: run
run: build
	docker run -p 8080:8080 -it mod_wasm:latest
