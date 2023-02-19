## Licoを導入する

### インストール
```bash
go install github.com/Hayao0819/lico@latest
```

### リポジトリを作成する

設定ファイルを含んだGitリポジトリを作成してください。

その後、リポジトリ内の設定ファイルとホームディレクトリでの関係を記述した`lico.list`を作成してください。

`lico.list`の書き方は[こちら](./config.md)を参照してください。



### リポジトリを登録し、適用する

リポジトリを作成したら、licoにリポジトリを登録します。

`set`コマンドでシンボリックリンクを作成します。

```bash
lico init <repository>
lico set
```
