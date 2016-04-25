BASE_BUILD_IMG = boostrap-golang
GO_DIR=/go/src/github.com/leocbs/main

base-build:
	docker build -t $(BASE_BUILD_IMG) .

build: base-build
	docker run --rm -v `pwd`:$(GO_DIR) -w $(GO_DIR) $(BASE_BUILD_IMG) go build

run: build
	./main
