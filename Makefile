.PHONY: all

all: clean get build

clean:
	rm -f drt

get:
	go get ./...

build:
	go build .

test: clean build
	./drt run samples/filehash/drt.yaml main.go
	./drt run samples/tensorflow/drt.yaml samples/tensorflow/tf_script.py

sample:
	cd samples/filehash && make

docker:
	docker build -t xebxeb/drt .
	docker run -it xebxeb/drt
