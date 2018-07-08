VANITY_DIR = opseeproto
EXAMPLE_DIR = examples
PROTO_DIR = proto
GIT_SHA := $(shell git describe --always HEAD)
DOCKER_IMAGE := "opsee/protobuf:$(GIT_SHA)"

all: VANITY_DIR EXAMPLE_DIR

VANITY_DIR: docker
	docker run --rm -it -v $$(pwd):/build $(DOCKER_IMAGE) /bin/bash -c 'cd /build/$(VANITY_DIR) && make generate'

docker:
	docker build -t $(DOCKER_IMAGE) .

EXAMPLE_DIR: docker
	docker run --rm -it -v $$(pwd):/go/src/github.com/opsee/protobuf $(DOCKER_IMAGE) /bin/bash -c 'cd /go/src/github.com/opsee/protobuf/$(EXAMPLE_DIR) && make generate && go test -v ./...'

clean:
	docker rmi $(DOCKER_IMAGE) || true
	$(MAKE) -C $(EXAMPLE_DIR) clean

.PHONY: docker clean all
