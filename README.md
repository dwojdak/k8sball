* clone repo
* go generate
* CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o k8sball .
* docker build
* docker run
