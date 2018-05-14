package main

import (
	"net/http"
	"log"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"encoding/json"
	"inj-init/utils"
	"inj-init/model"
)


func main() {
	http.HandleFunc("/startUp", StartUp)
	http.HandleFunc("/home", Home)
	err := http.ListenAndServe(":106", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("start server at 0.0.0.0:106")

}

func StartUp(w http.ResponseWriter, req *http.Request) {
	db,err := utils.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	var start model.StartUpModel
	db.Raw("select cover,`force`,number,msg,downloadUrl as download_url from btk_Version where client=?",req.URL.Query().Get("client")).First(&start)
	fmt.Println("start is",start)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "A Go Web Server")
	data, err := json.Marshal(start)
	if err != nil {
		log.Fatal("err get data: ", err)
	}

	w.Write(data)


}

func Home(w http.ResponseWriter, req *http.Request) {
	db,err := utils.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	var home model.Home
	//get banners
	var banners []model.Banner
	db.Raw("select id,title,type,cover,target_id,web_url from btk_Banner").Find(&banners)
	fmt.Println("banners is",banners)
	//get icons
	var icons []model.Icon
	db.Table("btk_Icon").Select("id, title,icon,target_id,web_url,type").Find(&icons)
	fmt.Println("icons is",icons)
	var cityToutiao model.CityToutiao
	cityToutiao.Icon = "http://inj-zone-img.bitekun.xin/resource/svg_life.png"
	cityToutiao.Cover = "http://inj-zone-img.bitekun.xin/resource/tt.jpg"
	cityToutiao.Title = "春夏新品大促 来啊来啊来啊"
	cityToutiao.Type = 1
	cityToutiao.TargetID = 1
	cityToutiao.WebURL = "bbb"
	fmt.Println("cityToutiao is",cityToutiao)
	//get zones
	var zones []model.ZoneItem
	db.Raw("SELECT ZoneID as zone_id, Name as name, Logo as logo, Brief as brief, MemberCount as member_count, Level as level, Tag as tag, CreateAt as create_at FROM btk_Zone WHERE Status = 0 ORDER BY CreateAt DESC").Limit(6).Offset(0).Find(&zones)
	fmt.Println("zones is",zones)
	//get goods 待开发
	//get activity
	var activities []model.Activity
    db.Raw("select tag,event_id,event_title,event_summary,event_thumb from t_event limit 0,4").Find(&activities)
    fmt.Println("activies is",activities)
    //get news
    var news []model.News
    db.Raw("select id,thumb,title,source,create_time,web_url from btk_News limit 0,6").Find(&news)
    fmt.Println("news is",news)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "A Go Web Server")
	home.Banners = banners
	home.Zones = zones
	home.Icons = icons
	home.CityTT = cityToutiao
	home.Activities = activities
	home.News = news

	data, err := json.Marshal(home)
	if err != nil {
		log.Fatal("err get data: ", err)
	}

	w.Write(data)
}