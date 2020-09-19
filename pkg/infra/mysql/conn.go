package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// Driver名
const driverName = "mysql"

// Conn 各repositoryで利用するDB接続(Connection)情報
type SqlHandler struct {
	Conn *sql.DB
}

// OriginalSqlHandlerにコネクションを格納
var OriginalSqlHandler *SqlHandler

// 接続するmysqlとのコネクションを生成
func init() {
	/* ===== データベースへ接続する. ===== */
	// DB設定

	var user string = os.Getenv("MYSQL_USER")
	var password string = os.Getenv("MYSQL_PASSWORD")
	var host string = os.Getenv("MYSQL_HOST")
	var port string = os.Getenv("MYSQL_PORT")
	var database string = os.Getenv("MYSQL_DATABASE")

	// 接続情報は以下のように指定する.
	// user:password@tcp(host:port)/database

	conn, err := sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}
	OriginalSqlHandler = new(SqlHandler)
	OriginalSqlHandler.Conn = conn
}

// トランザクション処理
func Transaction(db *sql.DB, txFunc func(*sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			log.Println("recover")
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Println("rollback")
			tx.Rollback()
		} else {
			log.Println("commit")
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}
