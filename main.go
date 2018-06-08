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
	"strconv"
)
var jsonf = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	CrawLine()
	//StartGetLineSegRelation()
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

	//fmt.Println(string(body))
	var data model.RouteData
	//
	jsonf.Unmarshal(body, &data)
	//fmt.Println("data is",data)

	fmt.Println("insert into `lines`(line_id,line_name) values(?,?)",data.RouteList[0].RouteID,data.RouteList[0].RouteName)

	for i:=0;i<len(data.RouteList) ;i++  {
       //db.Exec("insert into `lines`(line_id,line_name) values(?,?)",data.RouteList[i].RouteID,data.RouteList[i].RouteName)
       GetSeg(data.RouteList[i].RouteID,0)
       GetSeg(data.RouteList[i].RouteID,1)
	}
	fmt.Println("抓取完成!")


}

func GetSeg (line_id int64,updown int64) {
	id := strconv.FormatInt(line_id,10)
	up := strconv.FormatInt(updown,10)
	resp, err := http.Get("http://0523.bitekun.xin/getStation?line_id="+id+"&updown="+up)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("error")
	}
	fmt.Println(body)

}
