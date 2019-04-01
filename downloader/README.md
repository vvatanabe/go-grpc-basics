# file downloader using server-side streaming RPC

```
# build .proto
$ protoc -I ./downloader/proto --go_out=plugins=grpc:./downloader/proto/ ./downloader/proto/file.proto

# run server
$ go run ./downloader/server

# run client
$ go run ./downloader/client $FILE_NAME
```
