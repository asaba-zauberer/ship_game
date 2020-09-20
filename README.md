## 概要
<p>
プロトスプリントリーグでB班が制作した船のゲームサーバサイド<br>
API仕様書は SwaggerEditor に定義ファイルの内容を入力して参照してください。
</p>

SwaggerEditor: <https://editor.swagger.io> <br>
定義ファイル: `./api-document.yaml`<br>

※ Firefoxはブラウザ仕様により上記サイトからlocalhostへ向けた通信を許可していないので動作しません
- https://bugzilla.mozilla.org/show_bug.cgi?id=1488740
- https://bugzilla.mozilla.org/show_bug.cgi?id=903966

## DB構築
db/init/の1_ddl.sql, 2_dml.sqlを順に実行

## 環境変数
### API用のデータベースの接続情報を設定する
環境変数にデータベースの接続情報を設定します。<br>
ターミナルのセッション毎に設定したり、.bash_profileで設定を行います。

Macの場合
```
$ export MYSQL_USER=hoge \
    MYSQL_PASSWORD=hoge \
    MYSQL_HOST=127.0.0.1 \
    MYSQL_PORT=3306 \
    MYSQL_DATABASE=ca_hack
```

Windowsの場合
```
$ SET MYSQL_USER=hoge
$ SET MYSQL_PASSWORD=hoge
$ SET MYSQL_HOST=127.0.0.1
$ SET MYSQL_PORT=3306
$ SET MYSQL_DATABASE=ca_hack
```

## APIローカル起動方法
```
$ go run ./cmd/main.go
```

### ビルド方法
作成したAPIを実際にをサーバ上にデプロイする場合は、<br>
ビルドされたバイナリファイルを配置して起動することでデプロイを行います。
#### ローカルビルド
Macの場合
```
$ GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go
```

Windowsの場合
```
$ SET GOOS=linux
$ SET GOARCH=amd64
$ go build -o main ./cmd/main.go
```

このコマンドの実行で `main` という成果物を起動するバイナリファイルが生成されます。<br>
GOOS,GOARCHで「Linux用のビルド」を指定しています。
