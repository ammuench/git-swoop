SWOOPVERSION:= $(shell git describe --tags)
LDFLAGS += -X "main.swoopVersion=$(SWOOPVERSION)"
LDFLAGS += -X "main.goVersion=$(shell go version | sed -r 's/go version go(.*)\ .*/\1/')"


build:
	go build -ldflags '$(LDFLAGS)'

install:
	@echo "Installing git-swoop 🐦"
	go install -ldflags '$(LDFLAGS)'

