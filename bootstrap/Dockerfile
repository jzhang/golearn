FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/jzhang/golearn/bootstrap

ENV USER jzh
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET M07sTk-DqzuCJrAw

# Replace this with actual PostgreSQL DSN.
ENV DSN postgres://jzh@localhost:5432/bootstrap?sslmode=disable

WORKDIR /go/src/jzhang/golearn/bootstrap

RUN godep go build

EXPOSE 8888
CMD ./bootstrap