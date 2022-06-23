# Start from the latest golang base image
FROM golang:latest

WORKDIR /app
ENV APP_ENV dev

COPY go.mod go.sum ./
COPY . .
RUN go mod download

RUN go build ./cmd/faceit-task/main.go

EXPOSE 3001

CMD ["./main"]