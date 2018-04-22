FROM scratch
WORKDIR /
ADD k8sball /
# run the binary
CMD ["./k8sball"]
