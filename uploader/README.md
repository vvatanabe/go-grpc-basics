# file uploader using client-side streaming RPC

```
# build .proto
$ protoc -I ./uploader/proto --go_out=plugins=grpc:./uploader/proto/ ./uploader/proto/file.proto

# run server
$ go run ./uploader/server

# run client
$ go run ./uploader/client $FILE_NAME
```
