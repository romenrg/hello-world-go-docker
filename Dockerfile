FROM golang:1.8-alpine
ADD . /go/src/hello-world-go-docker
RUN go install hello-world-go-docker

FROM alpine:latest
COPY --from=0 /go/bin/hello-world-go-docker .
ENV PORT 8080
CMD ["./hello-world-go-docker"]