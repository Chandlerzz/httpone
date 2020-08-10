package serviceSql

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
// import "github.com/Chandlerzz/goChandler/jsonObj"
import "log"

type Idioms struct {
  DB *sql.DB
  Tx *sql.Tx
  Sentence string
  Result *sql.Rows
	Items []struct {
    Id int //`json:"id"`
		Name string //`json:"name"`
    Detail string //`json:"detail"`
    Comment string //`json:"comment"`

	} //`json:"items"`
}

// query idioms limit 5 randomly
func (idi *Idioms)Query(){
  rows, err := idi.DB.Query(idi.Sentence)
  if err != nil{
    log.Fatal(err)
  }
  idi.Result = rows
}
//insert one item
func (idi *Idioms)Insert(){
  stmt, err := idi.Tx.Prepare(idi.Sentence)
  if err != nil {
        log.Fatal(err)
      }
      _, err = stmt.Exec(idi.Items[0].Name,idi.Items[0].Detail,"")
      if err != nil {
              log.Fatal(err)
      }
  idi.Tx.Commit()
}

func(idi *Idioms)Update(){

}
func(idi *Idioms)Delete(){

}
