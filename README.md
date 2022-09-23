# 📝 Description

> A REST API boilerplate

# 🎉 Features

- TODOS: CRUD around TODOS

# 🧰 Installation

## Prerequisites

- go1.16
- Docker

Install go packages and dependencies before continue

```bash
go install
```

## Setting up PostgreSQL database

- This is will make a new PostgreSQL running in the standard port `5432`
- Please shutdown any previous conflicting PostgreSQL instances before starting
  this

```bash
docker-compose up -d
```

Check the database is up

```bash
docker logs -f content_pg
```

Check that you can log into a database with `psql`

```bash
docker exec -it content_pg psql -U content_pg_user content_pg_db
```

View tables

```psql
\dt
```

## Creating the initial database

Run initial migrations to set up initial database tables

```bash
yarn typeorm migration:run
```

# ⌨ Development

## ⚙ Running the app

```bash
go run main.go

# or air with live reload
air
```

## 🎮 Playground

After the application starts, go to `http://localhost:$PORT/swagger/index.html` to access swagger
playground

Observartion: You must change `$PORT` for the port to be used in your
environment

## 🧪 Running tests

Creating tests database Only integration tests are supported. Backend is spun up
on a special database

Tests use their own database. To create it:

```bash
docker exec -it content_pg psql -U content_pg_user -c "create database content_pg_db_test" content_pg_db
```

Note that in backend/config/typeorm.config.ts the content_db_test database is
configured to synchronize TypeORM migrations automatically, unlike the
development database.

```bash
# unit tests
go test
```

# 📦 Building

Before building application to production, make sure environment variables are
applied correctly

Building for production

```bash
go build
```

# ✅ TODO

- Add environment variables
- Add migrations to mongo-db-driver
- Add optional parameters to swagger
- Add bearer auth to swagger
- Add unit tests
- Add integration tests
- Add e2e tests
- Add log system
- Add health check

Observation: Some TODOS are spread across the code and need to be fixed ASAP

# 🛠 Built with

- [go](https://go.dev) - Go is an open source programming language supported by Google

# 👷 Authors

- [**Lucas Silva**](https://github.com/lucasmonstro) joao.galiano.silva@gmail.com -
  Developer

See also the list of [contributors](../../graphs/contributors) who participated
in this project

# 📝 License

Copyright © 2020 [**Lucas Silva**](https://github.com/lucasmonstro)  
This project is [MIT](https://opensource.org/licenses/MIT) licensed
