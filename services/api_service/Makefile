network:
	docker network create panshee-network

dserver:
	docker run --name api_service --network panshee-network -p 8080:8080 -e PANSHEE_API_SERVICE_DB_HOST=panshee_database -e GIN_MODE=release api_service:latest

dpostgres:
	docker run --name panshee_database --network panshee-network -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:latest

server:
	go run .

proto:
	rm -f api/panshee/v1/pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=api --go_out=./api/panshee/v1/pb \
	--go-grpc_out=./api/panshee/v1/pb \
	--grpc-gateway_out=./api/panshee/v1/pb \
	--openapiv2_out=./doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=api_service \
	./api/panshee/v1/proto/*.proto

protov2:
	protoc --proto_path=api --go_out=./api/panshee/v1/pb --go-grpc_out=./api/panshee/v1/pb --grpc-gateway_out=./api/panshee/v1/pb ./api/panshee/v1/proto/*.proto

linter:
	api-linter -I api/panshee/v1/proto/*.proto -o linter

docker:
	docker image rm hudyweas/panshee-account-service:latest
	docker build -t hudyweas/panshee-account-service:latest .
	docker push hudyweas/panshee-account-service:latest

.PHONY: proto server linter
