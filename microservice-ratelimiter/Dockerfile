FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./main.go

# Path: Dockerfile

FROM scratch
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

ENTRYPOINT ["./main"]