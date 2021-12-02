FROM golang:latest

COPY . /go/src
WORKDIR /go/src
RUN make install
ENTRYPOINT ["tail", "-f", "/dev/null"]