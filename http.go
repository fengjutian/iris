package main


import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
)

func httpPost() {
	resp, err := http.Post("http://www.baidu.com", 
	"app;ocation/x-www-form-urlencoded",	
	strings.NewReader("name=zzr"))
	
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

func main() {
	httpPost()
}