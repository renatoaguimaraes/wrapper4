ARG GOLANG_VERSION=1.18

FROM golang:${GOLANG_VERSION} as builder

WORKDIR /build

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go build -a -o wrapper ./cmd/wrapper

FROM gcr.io/distroless/base:latest-amd64

WORKDIR /

LABEL org.opencontainers.image.source https://github.com/renatoaguimaraes/wrapper4-k8s-job-istio

COPY --from=builder /build/wrapper .

ENTRYPOINT [ "/wrapper" ]
