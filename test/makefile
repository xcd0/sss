BIN := test.uf2

all:
	go generate
	tinygo build -target=pico -o $(BIN) .

flash:
	tinygo flash -target=pico .

install:
	go mod tidy

