package crawler

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"crypto/sha256"
	"strings"
)

func CheckDiff() {
	resp, err := http.Get("http://galaxyheavyblow.web.fc2.com/")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	// log.Println(string(body)) // HTMLが取得できていることを確認する。ただしshiftJISなので文字化けしている

	utf8String, err := shiftJIS2UTF8(string(body))
	//log.Println(utf8String) // HTMLが文字化けしていないことを確認する。
	hash := convertIntoHash(utf8String)
	log.Println(hash) // TODO ハッシュのバイト配列を確認する。
}

// sha256でハッシュ化する
func convertIntoHash(stringHTML string) []byte {
	s := sha256.New()
	s.Write([]byte(stringHTML))
	return s.Sum(nil)
}

// 文字コードをShiftJISからUTF8に変換する
func shiftJIS2UTF8(str string) (string, error) {
	strReader := strings.NewReader(str)
	decodedReader := transform.NewReader(strReader, japanese.ShiftJIS.NewDecoder())
	decoded, err := ioutil.ReadAll(decodedReader)
	if err != nil {
		return "", err
	}
	return string(decoded), err
}