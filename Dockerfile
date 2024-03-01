FROM golang:1.22.0

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o bin/main ./cmd

ARG APP_PORT
ENV APP_PORT=${APP_PORT}

EXPOSE ${APP_PORT}

CMD ["/app/bin/main"]
