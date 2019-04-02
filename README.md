# WEB+DB PRESS Vol.110 特集2「［速習］gRPC」サンプルコード

技術評論社刊「[WEB+DB PRESS Vol.110](https://gihyo.jp/magazine/wdpress/archive/2019/vol110)」の特集2「［速習］gRPC」のサンプルコードです。

## 目次

実行方法は各READMEをご覧ください。

- [単項RPC ── 1リクエスト／1レスポンス](https://github.com/vvatanabe/go-grpc-basics/tree/master/echo)
- [サーバサイドストリーミングRPC ── 1リクエスト／多レスポンス](https://github.com/vvatanabe/go-grpc-basics/tree/master/downloader)
- [クライアントサイドストリーミングRPC ── 多リクエスト／1レスポンス](https://github.com/vvatanabe/go-grpc-basics/tree/master/uploader)
- [双方向ストリーミングRPC ── 多リクエスト／多レスポンス](https://github.com/vvatanabe/go-grpc-basics/tree/master/chat)
- [gRPCのインタセプタ](https://github.com/vvatanabe/go-grpc-basics/tree/master/interceptor)

## 動作環境

サンプルコードは次の環境で動作確認済みです。

- Go 1.12
- grpc-go 1.19.1
- protobuf 3.7.0
- protoc-gen-go 1.3.1

## ライセンス

サンプルコードはMITライセンスで配布しています。

http://opensource.org/licenses/mit-license.php

## お問い合わせ

不具合があった場合は[本書Webページ](https://gihyo.jp/magazine/wdpress/archive/2019/vol110)よりお問い合わせをお願いいたします。

## ご注意

本サンプルコード、特集の内容に基づく運用結果について、著者、ソフトウェアの開発元および提供元、株式会社技術評論社は一切の責任を負いかねますので、あらかじめご了承ください。