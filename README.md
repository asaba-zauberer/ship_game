## 概要
<p>
ハッカソンで制作したゲームのサーバサイド<br>
仕様はAPI仕様書としているので SwaggerEditor にfune_api.ymlの内容を入力して参照してください。
</p>

SwaggerEditor: <https://editor.swagger.io> <br>


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
    MYSQL_DATABASE=hoge
```

Windowsの場合
```
$ SET MYSQL_USER=hoge
$ SET MYSQL_PASSWORD=hoge
$ SET MYSQL_HOST=127.0.0.1
$ SET MYSQL_PORT=3306
$ SET MYSQL_DATABASE=hoge
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
