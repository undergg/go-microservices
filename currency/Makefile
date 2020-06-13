.PHONY: protos

run: 
	go build main.go && ./main

protos:
	protoc -I protos/currency/ protos/currency/currency.proto --go_out=plugins=grpc:protos/currency/