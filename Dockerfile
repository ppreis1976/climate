# Etapa de construção
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o climatezipcode ./cmd/api

FROM scratch
COPY --from=builder /app/climatezipcode /climatezipcode
ENTRYPOINT ["/climatezipcode"]