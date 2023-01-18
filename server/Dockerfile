FROM golang:1.19-alpine as builder

WORKDIR /usr/src/app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application code
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o api .

FROM gcr.io/distroless/static-debian11 as final

COPY --from=builder /usr/src/app/api /

EXPOSE 8080

CMD ["/api"]
