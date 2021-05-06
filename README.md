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