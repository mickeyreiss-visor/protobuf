# builder stage is necessary because golang image doesn't include unzip.
FROM alpine as builder

# Download and install protoc.
WORKDIR /root
ADD https://github.com/google/protobuf/releases/download/v3.6.0/protoc-3.6.0-linux-x86_64.zip .
RUN mkdir protoc && unzip -d protoc protoc-3.6.0-linux-x86_64

FROM golang:1.10

COPY --from=builder /root/protoc/bin/ /usr/local/bin
COPY --from=builder /root/protoc/include/ /usr/local/include

WORKDIR src/github.com/opsee/protobuf
COPY ./ .
  
RUN \
  go install ./vendor/github.com/gogo/protobuf/protoc-gen-gogo && \
  go install ./opseeproto && \
  go install ./plugin/... && \
  go install ./protoc-gen-gogoopsee

WORKDIR /go
