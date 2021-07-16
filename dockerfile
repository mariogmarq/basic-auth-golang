FROM golang:1.16-alpine AS builder

WORKDIR /usr/app
COPY . .

RUN apk add build-base
RUN CGO_ENABLED=1 go build -ldflags '-s -w -extldflags "-static"' -o /usr/app/appbin cmd/apiTest/main.go

FROM alpine:3.14
RUN apk --update add ca-certificates && \
    rm -rf /var/cache/apk/*

RUN adduser -D appuser
USER appuser

COPY --from=builder /usr/app/appbin /home/appuser/app
WORKDIR /home/appuser/
EXPOSE 8080:8080

CMD ["./app"]