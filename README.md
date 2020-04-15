# finance-tracker

App for tracking finances. This is mostly for me to experiment with Go.

## Table of Contents

- [Running locally](#running-locally)

### Running locally

Setup is easy with [docker](https://www.docker.com/get-started) and [compose](https://docs.docker.com/compose/).

Simply run the following command from the root directory to run locally.

```bash
docker-compose -f docker-compose.yml -f docker-compose.local.yml up -d app
```
