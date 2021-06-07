FROM golang:alpine AS builder
RUN apk add git
ARG TAG=v0.1.5
RUN GO111MODULE=on go get -v github.com/miko/webgock@${TAG}

FROM alpine
WORKDIR /gock
EXPOSE 8080
COPY example /gock/example
COPY --from=builder /go/bin/webgock /bin/webgock
CMD ["/bin/webgock", "-addr", ":8080"]
