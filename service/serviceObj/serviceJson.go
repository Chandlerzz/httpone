package serviceObj

import (
  "github.com/gin-gonic/gin"
  "httpone/service/jsonObj"
  "httpone/service/serviceSql"
  "log"
  "fmt"
)
var db = serviceSql.OpenSql()

func IndexJson(c *gin.Context){
    var items jsonObj.AllItems
    items.GetJsonData(fullIndexName)
    c.JSON(200,items)
}

func ProductDataJson(c *gin.Context){
    name:=c.Param("name")
    fullName:="json/"+name+".json"
    var data jsonObj.Data
    data.GetJsonData(fullName)
    c.JSON(200,data)
}

func IdiomsJson(c *gin.Context){
    word := c.Query("word")
    idi := serviceSql.Idioms{}
    idi.DB = db
    idi.Sentence =`select * from idiom order by rand() limit 5`
    if word != "" {
      idi.Sentence = `select * from idiom where name like '%`+word+`%'`
    }
    idi.Query()
    rows := idi.Result
    for rows.Next() {
      var id int
      var name string
      var detail string
      var comment string
      rows.Columns()
      err := rows.Scan(&id,&name,&detail,&comment)
      if err != nil{
        log.Fatal(err)
      }
      newitems := append(idi.Items,struct{
        Id int
        Name string
        Detail string
        Comment string
        }{
        Id:id,
        Name:name,
        Detail:detail,
        Comment:comment,
      })
      idi.Items = newitems
    }
    fmt.Println(rows)
    c.JSON(200,idi.Items)

}
