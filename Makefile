postgres:
	docker run --name komikan-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.3-alpine
createdb:
	docker exec -it komikan-db createdb --username=root --owner=root komikan
migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/komikan?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/komikan?sslmode=disable" -verbose down
sqlc:
	sqlc generate
server:
	go run main.go
proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
  --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt paths=source_relative \
	--experimental_allow_proto3_optional \
  proto/*.proto
evans:
	evans --host localhost --port 9090 -r repl

.PHONY: postgres createdb migrateup migratedown sqlc proto evans server
