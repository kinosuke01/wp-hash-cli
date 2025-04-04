# wp-hash-cli

WordPressのパスワードハッシュを扱うためのコマンドラインツール

## インストール

```bash
go install github.com/kinosuke01/wp-hash-cli@latest
```

## 使用方法

### パスワードのハッシュ化

```bash
wp-hash-cli hash -p "your-password"
```

### パスワードとハッシュの検証

```bash
wp-hash-cli verify -p "your-password" -h "$P$B..."
```

## オプション

### hash コマンド
- `-p`: ハッシュ化するパスワード（必須）

### verify コマンド
- `-p`: 検証するパスワード（必須）
- `-h`: 検証するハッシュ（必須）

## 例

```bash
# パスワードのハッシュ化
$ wp-hash-cli hash -p "123456"
パスワード <123456> のハッシュ結果: <$P$B...>

# パスワードとハッシュの検証
$ wp-hash-cli verify -p "123456" -h "$P$B..."
パスワード <123456> とハッシュ <$P$B...> の検証結果: true
``` 