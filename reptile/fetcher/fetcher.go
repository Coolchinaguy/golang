package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string)([]byte,error)  {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取原网页失败，err = ",err)
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil,
		fmt.Errorf("wrong status code: %d",resp.StatusCode)
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return all,nil
}
