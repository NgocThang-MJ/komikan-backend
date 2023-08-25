#Build stage
FROM golang:1.20-alpine3.17 as builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY db/migration ./db/migration

ENV DATABASE_URL=postgres://root:secret@komikan-db:5432/komikan?sslmode=disable
ENV ORIGIN_ALLOWED=https://komikan.org
EXPOSE 8080
CMD ["/app/main"]
