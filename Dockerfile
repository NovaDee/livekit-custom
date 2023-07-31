FROM golang:1.20-alpine as builder

ARG TARGETPLATFORM
ARG TARGETARCH
RUN echo building for "$TARGETPLATFORM"

WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

# Copy the go source
COPY cmd/ cmd/
COPY pkg/ pkg/
COPY test/ test/
COPY tools/ tools/
COPY version/ version/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH GO111MODULE=on go build -a -o livekit-server ./cmd/server

FROM alpine

COPY --from=builder /workspace/livekit-server /livekit-server

# Run the binary.
ENTRYPOINT ["/livekit-server"]
#命令示例：docker build --build-arg REPOSITORY=docker.hanweb.com/hanwebbase/livertc --build-arg TAG=v1.4.4-c.1 -t docker.hanweb.com/hanwebbase/livertc:v1.4.4-c.1 .
