create-migrate-file:
	migrate create -ext sql -dir . -seq $(name)

migrate-table-up:
	migrate -path . -database "postgresql://root:1234@localhost:5432/test?sslmode=disable" up

migrate-table-down:
	migrate -path . -database "postgresql://root:1234@localhost:5432/test?sslmode=disable" down

migrate-table-up-version:
	migrate -path . -database "postgresql://root:1234@localhost:5432/test?sslmode=disable" force $(version)