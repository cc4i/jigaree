
# Image URL to use all building/pushing
DOCKER_IMAGE ?= weather
DOCKER_TAG ?= $(shell git rev-parse --abbrev-ref HEAD)
# Docker Hub username / password
DOCKER_PASS =
DOCKER_USER =


all:	manager

# Compile protos
pb:
	protoc -I=../../pb --go_out=. \
		--go-grpc_out=. \
		../../pb/weather.proto

test:	fmt	vet pb
	go test ./... -v -cover


# Build manager binary
manager:	fmt	vet pb
	go build -o weather main.go

run:	fmt	vet pb
	go run ./main.go

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Build the docker image
docker-build:
	docker build . -t ${DOCKER_IMAGE}:${DOCKER_TAG}

# Push the docker image
docker-push:
	#echo ${DOCKER_PASS} | docker login --username ${DOCKER_USER} --password-stdin
	docker push ${DOCKER_IMAGE}:${DOCKER_TAG}

