.DEFAULT_GOAL := build

build:
	go get ./...
	go build -o build/eopkg-graph cmd/main.go

install:
	sudo mv build/eopkg-graph /usr/local/bin/

uninstall:
	sudo /usr/local/bin/eopkg-graph

clean:
	rm -r build/