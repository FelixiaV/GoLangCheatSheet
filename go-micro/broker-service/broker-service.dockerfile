# base go image

FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/alpine
RUN chmod +x /app/brokerApp

# BUild a tiny docker image

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]
