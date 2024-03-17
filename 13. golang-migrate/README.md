# Golang-Migration Documentation

## CLI Install
* [CLI Install](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

Install
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Create Migration Files & Migrate with CLI
* [CLI docs](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)

Create Migration Files(One up and one down)
```bash
migrate create -ext sql -dir migrations -seq migration_name
```

Migration With CLI(Up & Down migration with cli)
```bash
# database `POSTGRESQL_URL` link 
export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/example?sslmode=disable&search_path=public'

# migration up
migrate -database ${POSTGRESQL_URL} -path migrations up 2

# migration down
migrate -database ${POSTGRESQL_URL} -path migrations down 2

```

## Migrate with program Docs
* [Postgres Setup](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)
* [Migrate code](https://github.com/golang-migrate/migrate)





