postgres:
	docker run -d --name my-postgres -e POSTGRES_USER=user -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=klp -p 5432:5432 postgres:latest

pgstart:
	docker start my-postgres

server:
	go run .

rm:
	docker rm my-postgres

stop:
	docker stop my-postgres

exec:
	docker exec -it my-postgres bash

lp:
	docker run -d -p8000:8000 antonboom/lines-provider

redis:
	docker run --name my-redis -p 6379:6379 -d redis

gen:
	protoc -I=./proto --go_out=./ --go-grpc_out=./ ./proto/*.proto

tidy:
	go mod tidy