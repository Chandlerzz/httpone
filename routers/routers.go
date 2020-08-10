package routers

import (
	"github.com/gin-gonic/gin"
	 "httpone/service/serviceObj"
	 // "httpone/service/serviceSql"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

  r := gin.Default()
  r.LoadHTMLGlob("templates/*")
  r.Static("static","./static")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
  //show the data for all the item
  r.GET("/index", serviceObj.ShowItem)
  //product detail
  r.GET("product/:name",serviceObj.ShowProductDetail)
  //adise the item
  r.GET("/revise/:name",serviceObj.ReviseProductDetail)
  //get json
  r.GET("/data/:name",serviceObj.ProductDataJson)
  //get index json
  r.GET("/indexjson",serviceObj.IndexJson)
  //get idioms json
	r.GET("/idiomsjson",serviceObj.IdiomsJson)
	//the page of fill the data
  r.GET("/dataup",serviceObj.ProductDataHtml)
  //get data from client
  r.POST("/pproductdata",serviceObj.PostProductData)
  // get revised data from client
  r.POST("/previseproductdata",serviceObj.PostReviseProductData)
	// get revised name from client
  r.GET("/rename",serviceObj.PostReviseProductName)

	return r
}
