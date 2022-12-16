# query-monitoring

---
Service providing API layer which works according postgres `pg_stat_statements` extension and provides queries monitor API based on their execution time.  
Using service provided API you can filter queries by statements (`select`, `insert`, `update`, `delete`) and sort by execution time (`asc`, `desc`).
API specification can be found in `api/swagger.yaml`

> Expect a high load of API calls additionally implemented cache layer. The cache expiration time is configured from the
> configuration, which gets the value from `CACHE_EXPIRATION` environment variable.


> Additionally, service providing Todo API layer for use demonstration purposes,
which makes CRUD dummy queries in the database so that the DB Query API has something to work with.


### Dependencies

---

- postgresql db (make sure postgres `pg_stat_statements` extension installed`max_connections` and `pg_stat_statements.track` configured)
- redis

### Env

---
```dotenv
DB_HOST=localhost
DB_NAME=postgres
DB_PASSWORD=12345
DB_PORT=5432
DB_USER=postgres
SERVICE_PORT=8080
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_DB=0
CACHE_EXPIRATION=20
```

### Run

---
`docker-compose up`

Docker compose runs Postgres Redis and application containers
### Test

---

`go test ./...`

### Implementation Notes

- Directory structure following [standard go project layout](https://github.com/golang-standards/project-layout)
- For keeping `main.go` file clean and readable used [fx](https://github.com/uber-go/fx) dependencies injection app
- For generating mocks used [moq](https://github.com/matryer/moq)
- For making migrations clean and separately used [goose](https://github.com/pressly/goose)
- As logger used famous [zap](https://github.com/uber-go/zap)
- As ORM used [gorm](https://github.com/go-gorm/gorm)
- As http web framework used [fiber](https://github.com/gofiber/fiber)

