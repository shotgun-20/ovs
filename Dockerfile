FROM alpine:3.15 as ovsbuild

# Install openvswitch to get the ovs-ofctl binary
RUN apk add --update --no-cache openvswitch

FROM golang:1.17-alpine3.15 as gobuild

RUN apk add --update --no-cache git

#add the working directory
ADD . /root/go/src/github.com/biwwy0/ovs-exporter

ENV GOPATH=/root/go

#build the GO binary
RUN cd /root/go/src/github.com/biwwy0/ovs-exporter \
    && ls -ld * \
    && go mod init \
    && go get -d \
    && go build .

FROM alpine:3.15

MAINTAINER "LeanNet" <info@leannet.eu>

#add ovs-ofctl dependecies
RUN apk add --update --no-cache libcap-ng libssl1.1

#copy the ovs-ofctl binary
COPY --from=ovsbuild /usr/bin/ovs-ofctl /usr/bin/ovs-ofctl

#copy the complied ovs-exporter binary
COPY --from=gobuild /root/go/src/github.com/biwwy0/ovs-exporter/ovs-exporter ./

ENTRYPOINT ["./ovs-exporter"]
