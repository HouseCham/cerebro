FROM golang:1.21-bullseye as build

WORKDIR /app

RUN useradd -u 2002 cerebro_user

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN go build \
    -ldflags "-linkmode external -extldflags -static" \
    -tags netgo \
    -o cerebro-service ./cmd

###
FROM scratch

COPY --from=build /etc/passwd /etc/passwd

COPY --from=build /app/cerebro-service cerebro-service

COPY config.json /config.json

USER cerebro_user

EXPOSE 3000

CMD ["/cerebro-service"]