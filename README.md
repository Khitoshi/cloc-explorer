# cloc-explorer

`cloc-explorer` は、GitHub リポジトリのコード行数、コメント行数、空行数を解析するツールです。

## インストール

以下のコマンドを使用してインストールできます。

```sh
go get github.com/Khitoshi/cloc-explorer
```

## 使用方法

以下のコマンドを使用して `cloc-explorer` を実行します。

```sh
cloc-explorer [OPTIONS]
```

### オプション

```
  --match-repository= Match GitHubRepository name
  --match-branch=     Match Branch name
  -h, --help          Show this help message
```

### 実際の出力

```
name                           files          blank        comment           code
-------------------------------------------------------------------------
C++                              0              0              0              0
Rust                             0              0              0              0
Go                               7             72             25            314
C                                0              0              0              0

Total                            7             72             25            314
```
