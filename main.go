package main

import (
	"net/http"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"taizhou-line-crawer/utils"
	"taizhou-line-crawer/model"
	"io/ioutil"
	"github.com/json-iterator/go"
)
var jsonf = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	CrawLine()
	/*http.HandleFunc("/taizhou/start",CrawLine)
	err := http.ListenAndServe(":1067", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("start server at 0.0.0.0:1067")*/
    //test
}

func CrawLine () {
	fmt.Println("抓取开始")
	db,err := utils.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()

	//http://61.132.47.90:8998/BusService/Require_AllRouteData/?TimeStamp=123
	//先获取所有线路
	resp, err := http.Get("http://61.132.47.90:8998/BusService/Require_AllRouteData/?TimeStamp=123")
	if err != nil {
		// handle error
	}

	// defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	var data model.RouteData
	//
	jsonf.Unmarshal(body, &data)
	fmt.Println("data is",data)

	for i:=0;i<len(data.RouteList) ;i++  {
       db.Exec("insert into lines(line_id,line_name) values(?,?)",data.RouteList[i].RouteID,data.RouteList[i].RouteName)
	}
	fmt.Println("抓取完成!")


}
