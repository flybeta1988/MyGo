package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type HttpTest struct {

}

func main() {
	httpTest := HttpTest{}
	httpTest.test1()
}

func (h HttpTest) test1() {
	url := "http://demo.t.me:8081/DB/test.php"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
