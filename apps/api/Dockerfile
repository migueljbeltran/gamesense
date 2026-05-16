FROM golang:1.24-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/gamesense-api ./cmd/api

FROM alpine:3.21

WORKDIR /app

COPY --from=build /bin/gamesense-api /app/gamesense-api

EXPOSE 8000

CMD ["/app/gamesense-api"]
