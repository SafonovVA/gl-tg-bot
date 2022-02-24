FROM golang:alpine as builder
RUN mkdir /build
COPY . /build
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build /app/

WORKDIR /app
CMD ["./gl-tg-bot"]