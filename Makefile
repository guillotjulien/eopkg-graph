.DEFAULT_GOAL := build

build:
	go get ./...
	go build -o build/eopkg-deps cmd/main.go

install:
	sudo mv build/eopkg-deps /usr/local/bin/

uninstall:
	sudo /usr/local/bin/eopkg-deps

clean:
	rm -r build/