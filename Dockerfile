FROM golang:alpine AS builder
RUN apk add git
RUN GO111MODULE=on go get -v github.com/miko/webgock

FROM alpine
ENTRYPOINT /bin/webgock
WORKDIR /gock
EXPOSE 8080
COPY example /gock/example
COPY --from=builder /go/bin/webgock /bin/webgock
