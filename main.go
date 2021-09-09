package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//respon, err := http.Get("https://book.douban.com/")//
	respon, err := http.Get("http://www.chinanews.com/")//utf编码
	if err != nil {
		panic(err)
	}
	defer respon.Body.Close()
	if respon.StatusCode != http.StatusOK {
		fmt.Sprintf("error status code", respon.StatusCode)
	}
	bodyReader := bufio.NewReader(respon.Body)
	e := determinEncodeing(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	res, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", res)
}

func determinEncodeing(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("获取失败:%v", err)
		return unicode.UTF8
	}
	re,_,_ := charset.DetermineEncoding(bytes, "")
	return re
}
