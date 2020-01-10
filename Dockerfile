from golang:1.11-stretch as builder
workdir /go/src/github.com/nimakaviani/example-websocket-app
copy . .
run go build -o /out/server /go/src/github.com/nimakaviani/example-websocket-app/server.go

from ubuntu:18.04
copy --from=builder /out/server /server
cmd ["/server"]

