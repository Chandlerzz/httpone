package jsonObj

// import "io/ioutil"
// import "log"
// import "encoding/json"
// import "github.com/gin-gonic/gin"

type GetJsonFile interface{
GetJsonData(fullName string)
Struct2JsonFile(fullName string)
}

type Data struct{
	Name string `json:"name"`
	Picture string `json:"picture"`
	Parameters []string `json:"parameters"`
	Characters []string `json:"characters"`
	Manual []string	`json:"manual"`
	Warnings []string `json:"warnings"`
}

func (s *Data)GetJsonData(pathName string){
  getJsonData(pathName,s)

}
func (s *Data) Struct2JsonFile(pathName string){
	struct2JsonFile(pathName,s)
}
