.PHONY: build

ifdef VERSION
    VERSION=${VERSION}
else
    VERSION="v0.0.1"
endif

build:
	GOOS=darwin go build -ldflags "-X github.com/dellkeji/helm-bcstest/pkg/version.Version=${VERSION}"  -o bin/darwin/amd64/bcstest main.go