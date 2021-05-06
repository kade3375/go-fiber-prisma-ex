.PHONY: start test dbSync dbGenerate

start:
	go run main.go

test:
	go test

dbSync:
	prisma-client-go db push

dbGenerate:
	prisma-client-go generate

#마이그레이션 실행
dbMigration:
	prisma-client-go migrate dev

#마이그레이션만 생성한다
dbGenerateMigration:
	prisma-client-go migrate dev --create-only