FROM golang:1.8.5-jessie as builder
WORKDIR /
ADD . /
# install xz
RUN apt-get update && apt-get install -y \
    xz-utils \
    go-bindata \
&& rm -rf /var/lib/apt/lists/*
# install UPX
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx
RUN go generate 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o k8sball ./main.go
# strip and compress the binary
RUN strip --strip-unneeded k8sball
RUN upx k8sball


FROM scratch
EXPOSE 3000
WORKDIR /
COPY --from=builder /* /
ADD k8sball /
CMD ["./k8sball"]
