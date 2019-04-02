# サーバサイドストリーミングRPC ── 1リクエスト／多レスポンス

サーバサイドストリーミングRPCを使用した、ファイルのダウンロード機能の例です。

## .protoファイルをコンパイルする

```
$ protoc -I ./downloader/proto --go_out=plugins=grpc:./downloader/proto/ ./downloader/proto/file.proto
```

## gRPCサーバを起動する

```
$ go run ./downloader/server
```

## gRPCクライアントプログラムを実行する

引数にダウンロードするファイル名を与えて実行します。

```
$ go run ./downloader/client $FILE_NAME
```