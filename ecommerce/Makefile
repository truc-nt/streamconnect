migrate-schema:
	jet -dsn="postgresql://root:1234@localhost:5432/test?sslmode=disable" -path=./internal/repositories -ignore-tables=schema_migrations

generate-wire:
	wire ./cmd