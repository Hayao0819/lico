## インストール

### makeを用いる

事前に`git`, `go`, `goreleaser`, `make`がインストールされている必要があります。

```bash
git clone https://github.com/Hayao0819/lico.git
cd lico
sudo make install
```

### Go Modules でインストール

事前に`go`がインストールされている必要があります。

```bash
go install github.com/Hayao0819/lico@HEAD
```

### バイナリをダウンロードする

#### HomeBrewの場合

```bash
brew tap Hayao0819/tap
brew install lico
```
#### Linux , macOSの場合

このスクリプトはWindowsでは動作しません

```bash
sudo sh -c "$(curl -L https://raw.githubusercontent.com/Hayao0819/lico/master/dl.sh)"
```

#### Windowsの場合

現在準備中


