network:
	docker network create panshee-network

dserver:
	docker run --name email_service --network panshee-network -p 8080:8080 -e GIN_MODE=release account_service:latest

server:
	go run .

proto:
	rm -f api/panshee/v1/pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=api --go_out=api/panshee/v1/pb \
	--go-grpc_out=api/panshee/v1/pb \
	--grpc-gateway_out=api/panshee/v1/pb \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=account_service \
	./api/panshee/v1/proto/*.proto

linter:
	api-linter -I api/panshee/v1/proto/*.proto -o linter

docker:
	docker build -t hudyweas/panshee-email-service:latest .

.PHONY: proto server linter
