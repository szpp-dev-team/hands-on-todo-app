# hands-on-todo-app

## 必要なツール

- docker
- go
- buf-cli

## 起動

以下のコマンドを実行すると backend server や db が同時に立ち上がります。

```shell
$ docker compose up -d --build
```

## 動作確認

[grpcui](https://github.com/fullstorydev/grpcui) を使うと GUI で gRPC の各メソッドを動作確認することが可能です。

```shell
$ grpcui -plaintext localhost:50051
```

https://github.com/szpp-dev-team/szpp-judge/assets/43411965/6d185773-11db-415d-9d71-8a92ae0762f5

## 停止・削除

```shell
$ docker compose down
```
