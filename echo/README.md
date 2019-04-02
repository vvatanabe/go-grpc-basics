# 単項RPC ── 1リクエスト／1レスポンス

単項RPCを使用した、クライアントからの入力をそのまま返却するEchoサーバの例です。

## .protoファイルをコンパイルする

```
$ protoc -I ./echo/proto --go_out=plugins=grpc:./echo/proto/ ./echo/proto/echo.proto
```

## gRPCサーバを起動する

```
$ go run ./echo/server
```

## gRPCクライアントプログラムを実行する

引数に適当な文言を与えて実行します。

```
$ go run ./echo/client hello
```