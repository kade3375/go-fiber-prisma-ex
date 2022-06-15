# go-fiber-prisma-ex

# 실행

## 기본

```shell
# DB migration (스키마 DB 마이그레이션 and DB 스키마 정의된 파일 생성)
make dbSync

# 실행
make dev

# 빌드
make build

# swagger
open http://localhost:3000/swagger

# graphql
open http://localhost:3000/graphiql
# or using insomnia
```

## DB 마이그레이션
```shell

# 마이그레이션 파일을 생성한다 (조건이 맞으면 ALTER, 아니면 CREATE 생성)
make dbGenerateMigration

# 실제 마이그레이션을 실행 (디폴트 값이 없거나 새로운 관계가 생기면 CREATE DDL 생성됨)
# 기존 데이터가 있었던 경우 모두 클리어 될수 있음 ..
make dbMigration
```

## 테스트

```shell
make test
```

### 커버리지 측정

```shell
make test
```

---
1. prisma ORM setting and TEST [o]
   - [prisma-client-go](https://github.com/prisma/prisma-client-go)
2. fiber settign and TEST [x]
   - [github.com/gofiber/fiber](https://github.com/gofiber/fiber)
3. Swagger [x]
4. GraphQL [x]