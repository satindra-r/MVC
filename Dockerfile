FROM golang:latest AS builder
LABEL author="saturn"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /mvc ./cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/database ./database
COPY --from=builder /app/pkg/views ./pkg/views
COPY --from=builder /mvc .

CMD ["cd /app"]
CMD ["/app/mvc"]