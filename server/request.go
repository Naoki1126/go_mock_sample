package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func RequestGet(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return []byte("fatal")
	}

	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b)) // htmlをstringで取得
	return b
}
