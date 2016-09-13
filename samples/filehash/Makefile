.PHONY: all

all: build

DOCKER_IMAGE_PREFIX=xebxeb/filehash

build:
	docker build --no-cache=true -t $(DOCKER_IMAGE_PREFIX) -f Dockerfile .

test: build
	docker run --rm -it -v $$PWD:/pwd $(DOCKER_IMAGE_PREFIX) /pwd/get_hash.py

publish:
	docker push $(DOCKER_IMAGE_PREFIX)
