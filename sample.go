package main
import "database/sql"
import _ "github.com/prestodb/presto-go-client/presto"

dsn := "http://user@localhost:8080?catalog=default&schema=test"
db, err := sql.Open("presto", dsn)

