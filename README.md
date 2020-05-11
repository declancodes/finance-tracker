# finance-tracker

App for tracking finances. This is mostly for me to experiment with Go.

## Table of Contents

- [Running locally](#running-locally)
  - [Backend](#backend)
  - [Frontend](#frontend)
- [DB Migrations](#db-migrations)

### Running locally

#### Backend

Setup is easy with [docker](https://www.docker.com/get-started) and [compose](https://docs.docker.com/compose/).

Simply run the following command from the repository root directory to run locally.

```bash
docker-compose -f docker-compose.yml -f docker-compose.local.yml up -d app
```

#### Frontend

The frontend for the project lives in the `js` directory. In development mode, it runs on `webpack-dev-server` on port 3000.
To run locally, simply `cd` into it and run `npm start`. You'll then be able to view it by navigating to `localhost:3000`.

```bash
cd js/ && npm start
```

### DB Migrations

DB migrations are SQL-based and performed with [flyway](https://flywaydb.org/) using their [official docker image](https://hub.docker.com/r/flyway/flyway/).
