create_migrate_file:
	migrate create -ext sql -dir . -seq $(name)

migrate_up:
	migrate -path . -database "postgresql://root:1234@localhost:5432/test?sslmode=disable" up

migrate_down:
	migrate -path . -database "postgresql://root:1234@localhost:5432/test?sslmode=disable" down

migrate_version:
	migrate -path . -database "postgresql://root:1234@localhost:5432/test?sslmode=disable" force $(version)