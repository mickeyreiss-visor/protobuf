VANITY_DIR = opseeproto
EXAMPLE_DIR = examples
PROTO_DIR = proto
GIT_SHA := $(shell git describe --always HEAD)
DOCKER_IMAGE := opsee/protobuf:$(GIT_SHA)

all: VANITY_DIR EXAMPLE_DIR

VANITY_DIR: docker
	docker run --rm -it -v $$(pwd):/build --workdir /build/$(VANITY_DIR) $(DOCKER_IMAGE) make generate

docker:
	docker build -t $(DOCKER_IMAGE) .

EXAMPLE_DIR: docker
	docker run --rm -it -v $$(pwd):/go/src/github.com/opsee/protobuf --workdir /go/src/github.com/opsee/protobuf/$(EXAMPLE_DIR) $(DOCKER_IMAGE) make generate

test: EXAMPLE_DIR
	docker run --rm -it -v $$(pwd):/go/src/github.com/opsee/protobuf --workdir /go/src/github.com/opsee/protobuf/$(EXAMPLE_DIR) $(DOCKER_IMAGE) go test ./...

clean:
	docker rmi $(DOCKER_IMAGE) || true
	$(MAKE) -C $(EXAMPLE_DIR) clean

.PHONY: docker clean all
