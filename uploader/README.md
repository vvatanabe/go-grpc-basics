# クライアントサイドストリーミングRPC ── 多リクエスト／1レスポンス

クライアントサイドストリーミングRPCを使用した、ファイルのアップロード機能の例です。

## .protoファイルをコンパイルする

```
$ protoc -I ./uploader/proto --go_out=plugins=grpc:./uploader/proto/ ./uploader/proto/file.proto
```

## gRPCサーバを起動する

```
$ go run ./uploader/server
```

## gRPCクライアントプログラムを実行する

引数にアップロードするファイル名を与えて実行します。

```
$ go run ./uploader/client $FILE_NAME
```