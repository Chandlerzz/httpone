package serviceObj

import (
  "github.com/gin-gonic/gin"
  "httpone/service/jsonObj"
)

const fullIndexName = "index.json"

func ShowItem(c *gin.Context){
  var items jsonObj.AllItems
  items.GetJsonData(fullIndexName)
  c.HTML(200,"show-items.tmpl",gin.H{"items":items.Items})
}

func ShowProductDetail(c *gin.Context){
    name:=c.Param("name")
    fullName:="json/"+name+".json"
    var data jsonObj.Data
    data.GetJsonData(fullName)
    c.HTML(200,"product-detail.tmpl",data)
}

func ReviseProductDetail(c *gin.Context){
    name:=c.Param("name")
    c.HTML(200,"data-revise.tmpl",gin.H{"name":name})
}

func ProductDataHtml(c *gin.Context){
    c.HTML(200,"dataup.tmpl",gin.H{"title":"login",})
}
