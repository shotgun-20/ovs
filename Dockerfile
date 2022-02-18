FROM debian:bullseye-slim as ovsbuild

# Install openvswitch to get the ovs-ofctl binary
RUN apt-get update && apt-get install -y openvswitch-common

FROM golang:1.17-bullseye as gobuild

RUN apt-get update && apt-get install -y git

#add the working directory
ADD . /root/go/src/github.com/shotgun-20/ovs-exporter

ENV GOPATH=/root/go

#build the GO binary
RUN cd /root/go/src/github.com/shotgun-20/ovs-exporter \
    && ls -ld * \
    && go mod init \
    && go get -d \
    && go build .

FROM debian:bullseye-slim

#add ovs-ofctl dependecies
RUN apt-get update \
    && apt-get install -y libcap-ng0 libssl1.1

#copy the ovs-ofctl binary
COPY --from=ovsbuild /usr/bin/ovs-ofctl /usr/bin/ovs-ofctl

#copy the complied ovs-exporter binary
COPY --from=gobuild /root/go/src/github.com/shotgun-20/ovs-exporter/ovs-exporter ./

ENTRYPOINT ["./ovs-exporter"]
