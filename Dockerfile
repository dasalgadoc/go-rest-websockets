ARG GO_VERSION=1.20.2

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk add --no-cache git

RUN apk --no-cache add ca-certificates && update-ca-certificates

RUN go env -w GOPROXY=direct

WORKDIR /src

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /go-rest-websockets .

FROM scratch AS runner

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /ect/ssl/certs/

COPY .env ./
COPY --from=builder /go-rest-websockets /go-rest-websockets

EXPOSE 8081

ENTRYPOINT ["/go-rest-websockets"]