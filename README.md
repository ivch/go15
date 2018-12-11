# 15 puzzle
Golang implementation of 15 puzzle with simple command line interface<br/>
Made as test task for the interview

## Godep
- `go get github.com/tools/godep`
- Save deps: `godep save -t ./...`
- Restore deps: `godep restore`

## Run tests
- `go test ./...`

## gocov Installation
- `go get github.com/axw/gocov/gocov`
- `go get -u gopkg.in/matm/v1/gocov-html`

## Calculate coverage
- `gocov test . | gocov-html > coverage.html`

## Run application
- restore dependecies
- `cd cli && go build -o app`
- `./app`

## Build Linux binary on any platform with docker
- run `make`
- `app` binary is placed in the `/cli` 

Due to specific input capturing libray can't be run in docker :(<br/
Can be extended to use http server. In that case can be run in docker
