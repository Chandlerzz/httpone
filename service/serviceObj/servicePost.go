package serviceObj

import (
  "github.com/gin-gonic/gin"
  "httpone/service/jsonObj"
  "log"
  "io/ioutil"
  "encoding/json"
  "fmt"
  "os"
)


func PostProductData(c *gin.Context){
  var data jsonObj.Data
  var items jsonObj.AllItems
  //get the byte from client
  json1,err1 :=c.GetRawData()
  if err1!=nil{
    log.Fatal(err1)
  }
  //json to object
  err := json.Unmarshal(json1, &data)
  if err != nil {
      panic(err)
  }
  item:=jsonObj.Item{Name:data.Name,Picture:data.Picture}
  items.GetJsonData(fullIndexName)
  items.Items = append(items.Items,item)
  fmt.Println(items)
  items.Struct2JsonFile(fullIndexName)
  filename:="json/"+data.Name+".json"
  err2 := ioutil.WriteFile(filename,json1,0644)
  if err2!=nil{
    log.Fatal(err2)
  }
}

func PostReviseProductData(c *gin.Context){
  var data jsonObj.Data
  //get the byte from client
  json1,err1 :=c.GetRawData()
  if err1!=nil{
    log.Fatal(err1)
  }
  //json to object
  err := json.Unmarshal(json1, &data)
  if err != nil {
      panic(err)
  }
  filename:="json/"+data.Name+".json"
  err2 := ioutil.WriteFile(filename,json1,0644)
  if err2!=nil{
    log.Fatal(err2)
  }
}

func PostReviseProductName(c *gin.Context){
  r := gin.Default()
  var data jsonObj.Data
  var items jsonObj.AllItems
  name := c.Query("name")
  rename := c.Query("rename")
  items.GetJsonData(fullIndexName)
  for k,_ := range items.Items{
    if (items.Items[k].Name == name) {
      items.Items[k].Name=rename
    }
  }
  fmt.Println(items.Items)
  items.Struct2JsonFile(fullIndexName)
  fileName:="json/"+name+".json"
  data.GetJsonData(fileName)
  data.Name = rename
  data.Struct2JsonFile(fileName)
  fileNewName := "json/"+rename+".json"
  err:=os.Rename(fileName,fileNewName)
  if err!=nil{
    log.Fatal(err)
  }
  c.Request.URL.Path="/index"
  r.HandleContext(c)
}
