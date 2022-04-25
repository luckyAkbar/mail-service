FROM golang:1.18-alpine

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

RUN rm -r src/

EXPOSE 5000

CMD ["./bin/main", "server"]