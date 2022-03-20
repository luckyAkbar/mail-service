FROM golang:1.17.1-alpine

WORKDIR /app

RUN mkdir src
RUN mkdir bin

WORKDIR /app/src

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o /app/bin main.go

WORKDIR /app

COPY .env .

RUN rm -r src/

EXPOSE 5000

CMD ["./bin/main", "server"]