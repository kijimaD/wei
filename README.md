# wei

Go template repository.

## install

```
$ go install github.com/kijimaD/wei@main
```

## docker run

```
$ docker run -v "$PWD/":/work -w /work --rm -it ghcr.io/kijimad/wei:latest
```

- 画像のビルドコマンド
  - wei build
  - こっちはカレントディレクトリのcsvを対象にすればいいのかな
- 記録機能
  - wei rec 55.55
  - グローバルで使いたい。保存場所を指定できるように。設定ファイル読み出し機能が必要

- config
  - csvの位置を指定できるようにする
