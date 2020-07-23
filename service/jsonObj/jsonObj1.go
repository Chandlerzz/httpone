package jsonObj

import "io/ioutil"
import "fmt"
import "log"
import "encoding/json"
//read json file
type AllItems struct{
	Items []Item `json:"items"`
}
type Item struct{
	Name string `json:"name"`
	Picture string `json:"picture"`
}


func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)

	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}

	for _, fi := range rd {
			println(fi.Name())
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
	}
	return s, nil
}


func (s *AllItems) GetJsonData(pathName string) {
	getJsonData(pathName,s)
}

func (s *AllItems) Struct2JsonFile(pathName string){
	struct2JsonFile(pathName,s)
}

func getJsonData(pathName string,s interface{}){
	json1,err:=ioutil.ReadFile(pathName)
  if err!=nil{
    log.Fatal(err)
  }
  err1:=json.Unmarshal(json1,s)
  if err1!=nil{
    log.Fatal(err1)
  }
}

func struct2JsonFile(pathName string,s interface{}){
	JSON,err:=json.Marshal(s)
	if err!=nil{
		log.Fatal(err)
	}
	err1:=ioutil.WriteFile(pathName,JSON,0644)
	if err1!=nil{
		log.Fatal(err1)
	}
}
