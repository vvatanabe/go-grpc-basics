# echo server using unary RPC

```
# build .proto
$ protoc -I ./echo/proto --go_out=plugins=grpc:./echo/proto/ ./echo/proto/echo.proto

# run server
$ go run ./echo/server

# run client
$ go run ./echo/client hello
```
