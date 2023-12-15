
start:
	docker compose up -d
	sleep 2
	migrate -path ./migrations -database "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable" up

stop:
	migrate -path ./migrations -database "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable" down
	docker compose down

restart: stop start

up:
	docker compose up -d
down:
	docker compose down
swag:
	swag init -g ./cmd/main.go


build:
	@go build -o .bin/app.exe cmd/store/main.go
run: build
	@.bin/app.exe
make migrate:
	migrate create -ext sql -dir migrations add_users_table

# Вместо qwerty введите действующий пароль от БД
migrate_up:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
migrate_down:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down