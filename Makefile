all:
	docker run --rm -v "$(CURDIR)":/go/src/github.com/ivch/go15  \
    	-w /go/src/github.com/ivch/go15 golang:1.10.3 sh -c 'make build'

build:
	go get github.com/tools/godep
	godep restore -v
	cd cli && CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o app