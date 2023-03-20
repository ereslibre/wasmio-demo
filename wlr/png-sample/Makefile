CFLAGS=$(shell pkg-config --cflags libpng libpng16 zlib)
LDFLAGS=$(shell pkg-config --libs libpng libpng16 zlib)
png: main.c
	clang -DPNG_SIMPLIFIED_READ_SUPPORTED -DPNG_SIMPLIFIED_WRITE_SUPPORTED $(CFLAGS) $(LDFLAGS) -o $@ $<
