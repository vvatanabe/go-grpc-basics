# gRPCのインタセプタ

インタセプタを使用した例です。

## .protoファイルをコンパイルする

```
$ protoc -I ./interceptor/proto --go_out=plugins=grpc:./interceptor/proto/ ./interceptor/proto/echo.proto
```

## gRPCサーバを起動する

```
$ go run ./interceptor/server
```

## gRPCクライアントプログラムを実行する

引数に適当な文言を与えて実行します。

```
$ go run ./interceptor/client hello
```