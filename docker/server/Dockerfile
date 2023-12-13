FROM centos:7
LABEL authors="anson@junerain.dev"

RUN yum update -y && \
    yum install -y wget && \
    yum install -y gcc && \
    yum install -y glibc-static && \
    yum clean all

ENV GOLANG_VERSION 1.20
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN wget https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go$GOLANG_VERSION.linux-amd64.tar.gz && \
    rm go$GOLANG_VERSION.linux-amd64.tar.gz

# copy code
COPY ../.. /server

WORKDIR /server/cmd

RUN go build -o main .

CMD ["./main"]
