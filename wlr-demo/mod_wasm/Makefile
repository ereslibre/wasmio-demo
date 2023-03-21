.PHONY: build
build:
	docker build -t mod_wasm:latest .

.PHONY: run
run: build
	docker run -p 8080:8080 --name mod_wasm-uuid -it mod_wasm:latest

.PHONY: clean
clean:
	docker rm -f mod_wasm-uuid
