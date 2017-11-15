FROM golang:1.9.2-alpine3.6 AS builder
RUN apk add -U git
WORKDIR /go/src/github.com/flexzuu/gas-price-server/
COPY . .

RUN go-wrapper download
RUN go-wrapper install

RUN go build

FROM alpine:3.6
WORKDIR /root/
COPY --from=builder /go/src/github.com/flexzuu/gas-price-server/gas-price-server .

EXPOSE 8080

CMD ["./gas-price-server"]