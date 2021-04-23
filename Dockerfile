FROM golang:alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" main.go

FROM scratch
COPY --from=builder /app/main /
CMD ["/main"]
