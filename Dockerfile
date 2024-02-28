FROM golang:1.22.0

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o bin/main ./cmd

EXPOSE 8080

CMD ["/app/bin/main"]
