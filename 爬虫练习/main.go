package main

import (
	"fmt"
	"io/ioutil"

	//"io/ioutil"
	"os"
)

func main() {
	//content, err := http.Get("http://tieba.baidu.com/f?kw=%C4%CF%BE%A9%B9%A4%B3%CC%D1%A7%D4%BA&fr=ala0&tpl=5")
	//if err != nil {
	//	log.Printf("网页获取失败")
	//}
	//
	//body,err := ioutil.ReadAll(content.Body)
	//if err != nil {
	//	log.Printf("缓存失败")
	//}
	//
	//fmt.Println(string(body))

	dir,err := ioutil.ReadDir("../video-server")
	if err != nil {
		fmt.Println("query dir wrong")
	}

	//io.Copy()

	file,err := os.Open("../reptile")
	file2,err := os.OpenFile("../reptile",32,0064)

	fmt.Println(file2)
	fmt.Println(file)

	//os.OpenFile()

	fmt.Println(dir)
}
