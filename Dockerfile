# Build Gccm in a stock Go builder container
FROM golang:1.12-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /go-ccmchain
RUN cd /go-ccmchain && make gccm

# Pull Gccm into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-ccmchain/build/bin/gccm /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["gccm"]
