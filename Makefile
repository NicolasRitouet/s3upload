BUILD_DIR?=build
REPO := s3upload
USER := NicolasRitouet
VERSION := "v0.1.0"

build:
	mkdir -p $(BUILD_DIR)/$(GOOS)
	GOOS=linux go build -o $(BUILD_DIR)/$(GOOS)/s3upload -ldflags "-X main.build `git rev-parse --short HEAD`" main.go commands.go version.go

clean:
	rm -rf build

install:
	go get github.com/codegangsta/cli
	go get github.com/goamz/goamz/aws
	go get github.com/goamz/goamz/s3

test:
	go test
