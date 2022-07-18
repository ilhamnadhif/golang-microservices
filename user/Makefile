
gen:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. model/proto/user.proto

run:
	go run main.go

build:
	go build main.go