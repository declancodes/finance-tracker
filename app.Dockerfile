FROM golang:1.14-alpine

WORKDIR /app

COPY . .

RUN go mod download && \
    go build -o tracker .

ENTRYPOINT [ "./tracker" ]