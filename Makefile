.PHONY: all build-binary build-image clean

all: build-binary build-image

build-binary:
	@echo "building apiserver binary"
	go build -o apiserver

build-image:
	@echo "building image"
	docker build -t pranganmajumder/go-basic-restapi:1.0.5 .

clean:
	rm -f apiserver