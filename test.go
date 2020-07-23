package main

import "github.com/gin-gonic/gin"
import "fmt"
import "httpone/foo"
import "log"
import "io/ioutil"
import "encoding/json"
import "httpone/service/jsonObj"
import "os"
// import "net/http"



func main() {
	fullIndexName:="index.json"
	foo.Bar()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("static","./static")
	//show the data for all the item
	r.GET("/index", func(c *gin.Context) {
		var items jsonObj.AllItems
		items.GetJsonData(fullIndexName)
		c.HTML(200,"show-items.tmpl",gin.H{"items":items.Items})

	})
//product detail
	r.GET("product/:name",func(c *gin.Context){
		name:=c.Param("name")
		fullName:="json/"+name+".json"
	  var data jsonObj.Data
		data.GetJsonData(fullName)
		c.HTML(200,"product-detail.tmpl",data)
	})
	//adise the item
	r.GET("/revise/:name",func(c *gin.Context){
		name:=c.Param("name")
		c.HTML(200,"data-revise.tmpl",gin.H{"name":name})
	})
	//get json
	r.GET("/data/:name",func(c *gin.Context){
		name:=c.Param("name")
		fullName:="json/"+name+".json"
		var data jsonObj.Data
		data.GetJsonData(fullName)
		c.JSON(200,data)
	})
	//get index.json
	r.GET("/indexjson",func(c *gin.Context){
		var items jsonObj.AllItems
		items.GetJsonData(fullIndexName)
		c.JSON(200,items)

	})
	//the page of fill the data

	r.GET("/dataup",func(c *gin.Context){
		c.HTML(200,"dataup.tmpl",gin.H{"title":"login",})
})




//get data from client
	r.POST("/login",func(c *gin.Context){
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
	})

	r.POST("/revise",func(c *gin.Context){
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
	})

	r.GET("/rename",func(c *gin.Context){
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
})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
