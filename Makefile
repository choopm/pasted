# Build for linux:
#   make xbuild-docker PLATFORM=Linux
# Build for osx:
#   make xbuild-docker PLATFORM=Darwin
# Build for windows
#   make xbuild-docker PLATFORM=windows/amd64

.PHONY: xbuild-docker clean

TARGET=bin
PLATFORM=$(shell uname)

xbuild-docker:
	@echo "Building for ${PLATFORM}"
	docker build --file Dockerfile \
		--ssh default \
		--target ${TARGET} \
		--platform ${PLATFORM} \
		--output dist/${PLATFORM} .

clean:
	rm -rf dist
