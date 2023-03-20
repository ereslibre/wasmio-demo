CFLAGS=$(shell pkg-config --cflags uuid)
LDFLAGS=$(shell pkg-config --libs uuid)

all: uuid.wasm uuid-cgi.wasm uuid-wws.wasm

uuid.wasm: main.c
	clang $(CFLAGS) $(LDFLAGS) -o $@ $<

uuid-cgi.wasm: main.c
	clang -DCGI $(CFLAGS) $(LDFLAGS) -o $@ $<

uuid-wws.wasm: main.c
	clang -DWWS $(CFLAGS) $(LDFLAGS) -o $@ $<

.PHONY: clean
clean:
	rm uuid.wasm uuid-cgi.wasm uuid-wws.wasm
