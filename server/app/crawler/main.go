package crawler

import (
	"io/ioutil"
	"log"
	"net/http"
	"crypto/sha256"
)

func CheckDiff() {
	resp, err := http.Get("http://galaxyheavyblow.web.fc2.com/")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body)) // HTMLが取得できていることを確認する

	hash := convertIntoHash(string(body))
	log.Println(string(hash[:])) // ハッシュ値を確認する
	// TODO 取得したHTMLとハッシュ値が文字化けしているので修正すること。
}

// sha256でハッシュ化する
func convertIntoHash(stringHTML string) [32]byte {
	return sha256.Sum256([]byte(stringHTML))
}