FROM golang as builder
WORKDIR /tmp/src/wxservice
ENV GO111MODULE on
ENV GOFLAGS="-mod=vendor"
ADD . /tmp/src/wxservice
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -v -o wxservice ./main

FROM alpine_with_lib
WORKDIR /root/
COPY --from=builder /tmp/src/wxservice/wxservice .
EXPOSE 7001
ENTRYPOINT  ["./wxservice"]