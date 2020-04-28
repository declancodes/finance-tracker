# finance-tracker

App for tracking finances. This is mostly for me to experiment with Go.

## Table of Contents

- [Running locally](#running-locally)
- [DB Migrations](#db-migrations)

### Running locally

Setup is easy with [docker](https://www.docker.com/get-started) and [compose](https://docs.docker.com/compose/).

Simply run the following command from the repository root directory to run locally.

```bash
docker-compose -f docker-compose.yml -f docker-compose.local.yml up -d app
```

### DB Migrations

DB migrations are SQL-based and performed with [flyway](https://flywaydb.org/) using their [official docker image](https://hub.docker.com/r/flyway/flyway/).
