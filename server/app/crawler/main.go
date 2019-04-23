package crawler

import (
	"io/ioutil"
	"log"
	"net/http"
)

func CheckDiff() {
	resp, err := http.Get("http://galaxyheavyblow.web.fc2.com/")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body)) // HTMLが取得できていることを確認する
}
