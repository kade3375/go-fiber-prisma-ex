.PHONY: start test dbSync dbGenerate

start:
	go run main.go

test:
	go test

dbSync:
	prisma-client-go db push

dbGenerate:
	prisma-client-go generate

dbMigration:
	prisma-client-go migrate dev --create-only