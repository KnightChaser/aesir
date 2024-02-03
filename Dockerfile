FROM golang:1.21.5-alpine3.18 AS gobuilder
ENV GO111MODULE=on \
    CGO_ENABLED=0
WORKDIR /build

# Keep all files and directories to the docker environment
COPY . .
COPY /datastructure ./datastructure
COPY /db ./db
COPY /web ./web
COPY /webRequestHandler ./webRequestHandler

RUN go mod download
RUN go build -o aesir main.go

FROM alpine:3.19.1
COPY --from=gobuilder /build .
EXPOSE 8080

CMD ["./aesir"]