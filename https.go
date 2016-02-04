package go_http_request
import (
	"net/http"
	"crypto/tls"
	"io/ioutil"
	"log"
)

func HttpsGet(url string) ([]byte,error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if resp==nil || resp.Body==nil || err!=nil {
		return nil,err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if body==nil || err!=nil {
		log.Println("2: ",err.Error())
		return nil,err
	}

	return body, nil
}

func HttpsGetInterface(header, url string, v interface{}) error {
	header = header + "HttpsGetInterface: "
	body,err := HttpsGet(url)
	if err != nil {
		return err
	}

	log.Println(header, url, string(body))
	return json.Unmarshal(body, &v)
}