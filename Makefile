build:
	@protoc -I Protos proto/sso/sso.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=path=source_relative
go run cmd/sso/main.go --config=./config/local.yaml