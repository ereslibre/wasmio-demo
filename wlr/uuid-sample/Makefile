CFLAGS=$(shell pkg-config --cflags uuid)
LDFLAGS=$(shell pkg-config --libs uuid)
uuid: main.c
	clang $(CFLAGS) $(LDFLAGS) -o $@ $<

.PHONY: clean
clean:
	rm uuid
