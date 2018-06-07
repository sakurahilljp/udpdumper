all: build


APP=udpdumper
VER = $(shell git describe --tags)
BUILDDATE=$(shell date '+%Y/%m/%d %H:%M:%S %Z')

LDFLAGS=-ldflags "-X main.Version=$(VER) -X \"main.BuildDate=$(BUILDDATE)\""

.PHONY: build install
build:
	go build $(LDFLAGS) -o $(APP)

install:
	go install $(LDFLAGS)

clean:
	rm -f $(APP)