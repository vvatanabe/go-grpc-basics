# 双方向ストリーミングRPC ── 多リクエスト／多レスポンス

双方向ストリーミングRPCを使用した、簡易的なチャット機能の例です。

## .protoファイルをコンパイルする

```
$ protoc -I ./chat/proto --go_out=plugins=grpc:./chat/proto/ ./chat/proto/chat.proto
```

## gRPCサーバを起動する

```
$ go run ./chat/server
```

## gRPCクライアントプログラムを実行する

引数に適当なユーザ名を与えて実行します。

```
$ go run ./chat/client $USER_NAME
```
