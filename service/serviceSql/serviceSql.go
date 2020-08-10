package serviceSql

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
// import "github.com/Chandlerzz/goChandler/jsonObj"
import "fmt"
import "log"

type sqlserv interface{
  Query()
  Insert()
  Delete()
  Update()
}


//opensql
func OpenSql() *sql.DB{
  db, err := sql.Open("mysql", "root:123456@/TEST?charset=utf8")
   if err!=nil{
     fmt.Println(err)
}
return db

}
// Prepare for insert items
func BeginSql(db *sql.DB) *sql.Tx{
  tx, err := db.Begin()
  if err != nil {
          log.Fatal(err)
  }
  return tx
}
