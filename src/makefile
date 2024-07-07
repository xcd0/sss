BIN := sss46g.uf2

all:
	go generate
	tinygo build -target=pico -o $(BIN) .

flash:
	go generate
	tinygo flash -target=pico .

flash-monitor:
	go generate
	tinygo flash -target=pico . -monitor

install:
	go mod download github.com/itchio/lzma
	go mod tidy

