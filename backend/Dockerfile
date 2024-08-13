FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o floodguard-backend
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/floodguard-backend .
ENTRYPOINT [ "./floodguard-backend" ]
