from golang:1.11-alpine3.8 as builder
workdir /go/src/github.com/nimakaviani/knative-sample-app
copy . .
run go build -o out/server github.com/nimakaviani/example-websocket-app/server.go

from alpine:3.8
copy --from=builder /go/src/github.com/nimakaviani/example-websocket-app/out/server /server
cmd ["/server"]

