db:
	docker run -d --name mysql-api -e MYSQL_PASSWORD=abc --env MYSQL_USER=abc --env MYSQL_DATABASE=api --env MYSQL_ROOT_PASSWORD=root -p 3306:3306 mysql:latest

start-db:
	docker start mysql-api

stop-db:
	docker stop mysql-api

run:
	go run ./cmd/main.go
