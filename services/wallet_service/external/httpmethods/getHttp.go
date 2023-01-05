package httpmethods

import (
	"encoding/json"
	"net/http"
)

func GetHttpJsonAndDecode(req string, v interface{}) (err error) {
	resp, err := http.Get(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&v)

	return
}
