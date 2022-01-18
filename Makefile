.PHONY: build

build:
	GOOS=darwin go build -o bin/darwin/amd64/teststack *.go