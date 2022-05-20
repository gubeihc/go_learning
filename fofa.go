package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Fofa struct {
	FOFAEmail string `json:"FOFA_email"`
	FOFAKey   string `json:"FOFA_key"`
}

var fofa Fofa

type Fofa_data struct {
	Error   bool       `json:"error"`
	Size    int        `json:"size"`
	Page    int        `json:"page"`
	Mode    string     `json:"mode"`
	Query   string     `json:"query"`
	Results [][]string `json:"results"`
}

func get_fofa(key string, size string) {
	//base64加密 key
	basekey := base64.StdEncoding.EncodeToString([]byte(key))
	//构造url
	url := fmt.Sprintf("https://fofa.info/api/v1/search/all?email=%s&key=%s&qbase64=%s&size=%s", fofa.FOFAEmail, fofa.FOFAKey, basekey, size)
	get, err := http.Get(url)
	if err != nil {
		log.Println("获取fofa信息失败:", err)
		return
	}
	defer get.Body.Close()
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		log.Println("获取fofa信息失败:", err)
		return
	}
	var fofa_data Fofa_data
	err = json.Unmarshal(all, &fofa_data)
	if fofa_data.Error == true {
		fmt.Println("请检测key 是否可以正常使用")
		return
	}
	log.Println("正在搜索:", key)
	if err != nil {
		log.Println("获取fofa信息失败:", err)
		return
	}
	for _, v := range fofa_data.Results {
		fmt.Println(v[0])
	}
}

//读取配置文件
func config_fofa() {
	open, err := os.Open("config_init.json")
	if err != nil {
		log.Fatal("读取配置文件错误")
	}
	defer open.Close()
	decoder := json.NewDecoder(open)
	err = decoder.Decode(&fofa)
	if err != nil {
		log.Fatal("解析配置文件错误")
	}

}
func main() {
	config_fofa()
	key := flag.String("key", "", "搜索关键词")
	size := flag.String("size", "10000", "搜索数量")
	flag.Parse()
	if *key == "" {
		flag.PrintDefaults()
		return
	}
	get_fofa(*key, *size)
}
