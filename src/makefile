

all:
	go generate
	tinygo build -target=pico -o `cat go.mod | grep module | awk '{print $$2}'`.uf2 .
flash:
	tinygo flash -target=pico .


install:
	go mod download github.com/itchio/lzma
	go mod tidy

