network:
	docker network create panshee-network

dserver:
	docker run --name account_service --network panshee-network -p 8080:8080 -e PANSHEE_ACCOUNT_SERVICE_DB_HOST=panshee_database -e GIN_MODE=release account_service:latest

dpostgres:
	docker run --name panshee_database --network panshee-network -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:latest

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
	rm -f libs/pbs/account_service/pb
	cp api/panshee/v1/pb/*.go ../../libs/pbs/account_service/pb

linter:
	api-linter -I api/panshee/v1/proto/*.proto -o linter

docker:
	docker build -t panshee-account-service:latest .

.PHONY: proto server linter
