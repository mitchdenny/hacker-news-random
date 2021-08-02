default: build

clean:
	rm -rf build

build: clean hnr

hnr:
	go build -o build/hnr cmd/hnr/main.go

install:
	cp build/hnr /usr/local/bin

run: build
	clear
	build/hnr