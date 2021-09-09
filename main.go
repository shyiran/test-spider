package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	respon,err:=http.Get("https://book.douban.com/")
	if err != nil {
		panic(err)
	}
	defer respon.Body.Close()
	if respon.StatusCode != http.StatusOK{
		fmt.Sprintf("error status code",respon.StatusCode)
	}
	res,err:=ioutil.ReadAll(respon.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s",res)
}
