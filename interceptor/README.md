# echo server using interceptor

```
# build .proto
$ protoc -I ./interceptor/proto --go_out=plugins=grpc:./interceptor/proto/ ./interceptor/proto/echo.proto

# run server
$ go run ./interceptor/server

# run client
$ go run ./interceptor/client hello
```
