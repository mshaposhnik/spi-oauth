FROM registry.access.redhat.com/ubi8/go-toolset:1.16.7-5 as builder
ENV GOPATH=/go/
USER root

WORKDIR /spi-oauth

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Copy the go source
COPY main.go main.go
COPY main_test.go main_test.go
COPY controllers/ controllers/

# build service
RUN export ARCH="$(uname -m)" && if [[ ${ARCH} == "x86_64" ]]; then export ARCH="amd64"; elif [[ ${ARCH} == "aarch64" ]]; then export ARCH="arm64"; fi && \
    go test -v ./... && \
    CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} GO111MODULE=on go build -a -o spi-oauth main.go

FROM registry.access.redhat.com/ubi8-minimal:8.7-1049.1675784874

COPY --from=builder /spi-oauth/spi-oauth /spi-oauth

WORKDIR /
USER 65532:65532

ENTRYPOINT ["/spi-service"]

