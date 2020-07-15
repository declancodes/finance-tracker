FROM golang:1.14

WORKDIR /app

COPY /app .

RUN go mod download && \
    go build -o tracker .

ENTRYPOINT [ "./tracker" ]
