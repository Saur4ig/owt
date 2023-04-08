FROM golang:alpine AS builder

ENV CGO_ENABLED=1

RUN apk add --no-cache \
    gcc \
    musl-dev

ARG APP_NAME=owt_challenge
WORKDIR /build-dir
COPY . .
RUN apk add --update gcc musl-dev
RUN go build -ldflags='-s -w -extldflags "-static"' -o /${APP_NAME} cmd/owt/main.go


FROM scratch
COPY --from=builder /${APP_NAME} /${APP_NAME}

EXPOSE 8080

ENTRYPOINT ["/owt_challenge"]