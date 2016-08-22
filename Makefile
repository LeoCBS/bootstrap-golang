BASE_BUILD_IMG = go-scratchpad
GO_DIR=/go/src/github.com/leocbs/main
RUN_GO=docker run -v `pwd`:$(GO_DIR) -w $(GO_DIR) $(BASE_BUILD_IMG) 

base-build:
	docker build -t $(BASE_BUILD_IMG) .

build: base-build
	$(RUN_GO) go build

run: build
	./main

run-pointers: base-build
	$(RUN_GO) go run pointers/pointers.go	
