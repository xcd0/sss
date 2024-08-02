BIN := sss46g.uf2

all:
	cat makefile
build:
	go generate
	tinygo build -target=pico -o $(BIN) .
flash:
	tinygo flash -target=pico .
install:
	go mod download github.com/itchio/lzma
	go mod tidy
clean:
	rm -rf *.uf2
