# query-monitoring

---
Service providing API layer which works according postgres `pg_stat_statements` extension and provides queries monitor API based on their execution time.  
You can filter queries by statements (`select`, `insert`, `update`, `delete`) and sort by execution time (`asc`, `desc`) via service provided API .
API specification can be found in the `api/swagger.yaml`

> I implemented a cashing layer as a protective measure in case of high load. The cache expiration time is configured from the
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
DB_HOST=postgres
DB_NAME=postgres
DB_PASSWORD=12345
DB_PORT=5432
DB_USER=postgres
SERVICE_PORT=8080
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_DB=0
CACHE_EXPIRATION=20
```

### Run

---
`docker-compose up`

Docker compose runs Postgres Redis and app container.

I have implemented migrations as an admin processes so, it runs on startup of the application. After successfull migration we should have following tables in database.

`postgres=# \d`
List of relations

| Schema  |          Name           |   Type   |  Owner   |
|---------|-------------------------|----------|----------|
| public  | goose_db_version        | table    | postgres |
| public  | goose_db_version_id_seq | sequence | postgres |
| public  | pg_stat_statements      | view     | postgres |
|  public | todos                   | table    | postgres |

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
- For caching used Redis and [redis go client](https://github.com/go-redis/redis)

