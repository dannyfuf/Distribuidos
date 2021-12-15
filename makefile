
proto:
	@(PATH="$(PATH):$(go env GOPATH)/bin" && cd protos && protoc --go_out=plugins=grpc:../servers/$(name) --go_opt=paths=source_relative $(name).proto)

server:
	@(cd servers && go run $(name)_server.go)

install:
	@apt update
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go get -u github.com/golang/protobuf/protoc-gen-go
	@apt install -y protobuf-compiler
	@(cd servers && go mod tidy)

clean_data:
	rm -rf servers/data