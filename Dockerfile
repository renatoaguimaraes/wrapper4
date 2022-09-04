ARG GOLANG_VERSION=1.18

FROM golang:${GOLANG_VERSION} as builder

WORKDIR /build

ENV GO111MODULE=on
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go build -a -o wrapper ./cmd/wrapper
RUN go build -buildmode=plugin -o istio-proxy-plugin.so ./cmd/plugin/istio-proxy

FROM gcr.io/distroless/base:latest-amd64

WORKDIR /

LABEL org.opencontainers.image.source https://github.com/renatoaguimaraes/wrapper4-k8s-job-istio

COPY --from=builder /build/wrapper .
COPY --from=builder /build/istio-proxy-plugin.so .

ENV WRAPPER_PLUGIN_PATH=/istio-proxy-plugin.so

ENTRYPOINT [ "/wrapper", "ls", "-l" ]
