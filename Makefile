# This Makefile provides targets for building, running, and cleaning the backend server.

.PHONY: all build run clean

all: build

build:
	go build -o backend-server cmd/delivery/main.go

run:
	./backend-server

clean:
	rm -f backend-server