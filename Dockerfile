FROM golang:1.16-alpine AS builder
WORKDIR /go/src/github.com/wzshiming/ssdb/
COPY . .
RUN go install ./cmd/...

FROM alpine
COPY --from=builder /go/bin/ssdb /usr/local/bin/
ENTRYPOINT /usr/local/bin/ssdb
