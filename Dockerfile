FROM alpine:latest
COPY gas-price-server bin

EXPOSE 8080

CMD ["gas-price-server"]
# CMD ["ls"]