# Build stage
FROM golang:1.25.9-alpine3.23 AS build
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Runtime stage
FROM alpine:3.23
WORKDIR /app
COPY --from=build /app/main .
COPY app.env .

EXPOSE 8080
CMD ["/app/main"]