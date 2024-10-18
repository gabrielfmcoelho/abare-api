FROM golang:1.23 AS builder
ARG CGO_ENABLED=0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server cmd/main.go

FROM scratch AS runner
COPY --from=builder /app/server /server
CMD ["/server"]