package controllers_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"gozen/controllers"
)

var (
	server *httptest.Server
	reader io.Reader
)

const (
	gozen = "gozen"
)

type commonData struct {
	method       string
	filePath     string
	requestPath  string
	exp, act     interface{}
	statusCode   int
	sessionToken string
}

func init() {
	server = httptest.NewServer(controllers.Routes())
}

func checkGozen(t *testing.T, category string, data commonData) {
	checkImpl(t, category, data)
}

func checkImpl(t *testing.T, category string, data commonData) {

	// 期待値のJSONをデコード
	file, err := os.Open(fmt.Sprintf("testdata/%s", data.filePath))
	if err != nil {
		t.Error(err)
	}
	decoder := json.NewDecoder(file)
	defer file.Close()

	expect := data.exp
	decoder.Decode(&expect)

	// HTTPリクエスト
	url := fmt.Sprintf("%s/api/%s/%s", server.URL, category, data.requestPath)
	t.Log(url)
	reader = strings.NewReader("")

	if data.method == "" {
		data.method = "GET" // 指定がなければGETとする
	}
	request, err := http.NewRequest(data.method, url, reader)

	if data.sessionToken != "" {
		request.Header.Add("X-Session-Token", data.sessionToken)
	}
	t.Log(request.Header)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	actual := data.act

	err = json.Unmarshal(body, &actual)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != data.statusCode {
		t.Errorf("%s: \nSuccess expected: %d, \nbut actual value: %d", data.requestPath, data.statusCode, resp.StatusCode)
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%s: \nSuccess expected: %s, \nbut actual value: %s", data.requestPath, expect, actual)
	}

}
