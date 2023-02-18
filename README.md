## Lico - シンプルで軽量なDotfilesマネージャー

Licoはシンプルで軽量なDotfilesマネージャーです。全てのファイルをGitで管理し、自由なファイル構造を維持できます。

このツールは[ssh0/dot](https://github.com/ssh0/dot)に影響を受けて開発されています。

## ドキュメント

### Licoを導入する

#### インストール
```bash
go install github.com/Hayao0819/lico@latest
```

#### リポジトリを作成する

設定ファイルを含んだGitリポジトリを作成してください。

その後、リポジトリ内の設定ファイルとホームディレクトリでの関係を記述した`lico.list`を作成してください。

`lico.list`の書き方は以下の通りです。

```txt
<リポジトリ内のパス>:<ホームディレクトリでのパス>
```

リストファイルではGolangのテンプレートを使用できます。テンプレートを用いてOSごとに分岐して設定ファイルを配置することができます。

```txt
zsh/{{ .OS }}/zshrc:{{ .Home }}/.zshrc
```

テンプレートで利用可能な変数の一覧は`env`コマンドで確認できます。

#### リポジトリを登録し、適用する

リポジトリを作成したら、licoにリポジトリを登録します。

```bash
lico init <repository>
lico set
```

## Todo

- Gitの操作を[go-git](https://github.com/go-git/go-git)で書き直す
- rmlinkコマンドを実装する
- rmファイルを実装する
- unlink時にrmlinkコマンドを実行する

## Special Thanks

- [まちかどまぞく](https://www.tbs.co.jp/anime/machikado/)
- [ssh0/dot](https://github.com/ssh0/dot)
- [mazen160/go-random](https://github.com/mazen160/go-random)
- [spf13/cobra](https://github.com/spf13/cobra)
- [watasuke102/mit-sushi-ware](https://github.com/watasuke102/mit-sushi-ware)
