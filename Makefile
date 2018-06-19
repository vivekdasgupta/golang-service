PROJECT = drmshows
OUT_DIR = temp
BUILD_DIR = build
AWS_KEYS = creds.yml

all: clean build
.PHONY: all

# install dependencies

# install dependencies
# Use this make target for source dependencies locally on desktop  ***
deps:
	echo "Installing dependencies ..."

.PHONY: deps


build:
	export GOPATH=$(GOPATH):$(PWD);  GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o application
.PHONY: build

build_zip:
	cp application temp/
	cd temp; zip drmshows *
.PHONY: build_zip

# Remove all build artifacts.
#
# Example:
#   make clean
clean:
	sudo rm -rf $(OUT_DIR)
	sudo rm -rf $(KEY_DIR)
	sudo rm -rf $(BUILD_DIR)
	sudo rm -rf tmp
.PHONY: clean
