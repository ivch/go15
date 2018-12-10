# 15 puzzle
Golang implementation of 15 puzzle with simple command line interface
Made as test task for the interview

## Godep
- `go get github.com/tools/godep`
- Save deps: `godep save -t ./...`
- Restore deps: `godep restore`

## Run tests
- `go test ./...`

## Run application
- restore dependecies
- `cd cmd && go build -o app`
- `./app`

Due to specific input capturing libray can't be run in docker :(
Can be extended to use http server. In that case can be run in docker