# simple chat using bidirectional streaming RPC

```
# build .proto
$ protoc -I ./chat/proto --go_out=plugins=grpc:./chat/proto/ ./chat/proto/chat.proto

# run server
$ go run ./chat/server

# run client
$ go run ./chat/client $USER_NAME
```
