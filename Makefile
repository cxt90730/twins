.PHONY: proto

proto:
	protoc -I proto/ proto/heartbeat.proto --go_out=plugins=grpc:proto

build:
	go build

all: proto build