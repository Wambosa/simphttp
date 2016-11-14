package simphttp

import (
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func GetString(link string) (string, error) {

	resp, err := http.Get(link)

	if err != nil {return "", err}

	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)

	if err != nil {return "", err}

	return string(buf), err
}

func GetJson(link string) (map[string]interface{}, error) {

	raw, err := GetString(link)

	if err != nil { return nil, err}

	var newMap map[string]interface{}

	json.Unmarshal([]byte(raw), &newMap)

	return newMap, err
}

func Query(endpoint string, queryF string, params ...string) (map[string]interface{}, error){
	
	escaped := make([]interface{}, len(params))
	
	for i, s := range params{
		escaped[i] = url.QueryEscape(s)
	}
	
	q:= fmt.Sprintf(queryF, escaped...)
	
	return GetJson(fmt.Sprintf("%s?%s", endpoint, q))
}