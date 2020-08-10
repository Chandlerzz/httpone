package main

import "httpone/routers"




func main() {
	routersInit := routers.InitRouter()
	routersInit.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
