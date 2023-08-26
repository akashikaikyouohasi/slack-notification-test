
## go

### メモ
`x4 := "test"     // 文字列で初期化。varも省略する記法`

CodeUriでgo file not foundとなったら以下を実施する必要があるかも
```
$ go mod init notification-slack
$ go mod tidy
```

### ファイルの説明
go.mod: モジュールのインポートパスとバージョン情報
go.sum: 依存先モジュールの記録ファイル

### モジュールインストール
`notification-test/hello-world`ディレクトリで実施

```
go get -u github.com/slack-go/slack
```

