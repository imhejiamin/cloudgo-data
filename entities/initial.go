package entities

import (
    "database/sql"
    "github.com/go-xorm/xorm"
    _ "github.com/go-sql-driver/mysql"
)

var mydb *sql.DB
var engine *xorm.Engine

func init() {
    //https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
    if err != nil {
        panic(err)
    }
    mydb = db

    en, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
    //tx, err := mydb.Begin()
    checkErr(err)
    engine = en
}

// SQLExecer interface for supporting sql.DB and sql.Tx to do sql statement
type SQLExecer interface {
    Exec(query string, args ...interface{}) (sql.Result, error)
    Prepare(query string) (*sql.Stmt, error)
    Query(query string, args ...interface{}) (*sql.Rows, error)
    QueryRow(query string, args ...interface{}) *sql.Row
}

// DaoSource Data Access Object Source
type DaoSource struct {
    // if DB, each statement execute sql with random conn.
    // if Tx, all statements use the same conn as the Tx's connection
    SQLExecer
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
