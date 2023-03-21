package helper

import (
	"encoding/json"
	"net/http"
)

// for decoding the body
func GetReqBody(r *http.Request, res interface{}) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(res)
	PanicIfError(err)
}

// for encoding the body for request
func WriteResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}
