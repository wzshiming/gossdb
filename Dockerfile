FROM golang:1.10-alpine3.8 AS builder
WORKDIR /go/src/github.com/wzshiming/ssdb/
COPY . .
RUN go install ./cmd/...

FROM scratch
COPY --from=builder /go/bin/ssdb /usr/local/bin/
ENTRYPOINT ssdb
