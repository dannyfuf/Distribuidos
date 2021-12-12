FROM golang:latest

COPY . /go/src
WORKDIR /go/src
RUN make install
RUN apt install vim
ENTRYPOINT ["tail", "-f", "/dev/null"]