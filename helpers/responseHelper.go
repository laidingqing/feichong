package helpers

import (
	"encoding/json"
	"net/http"
)

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func SetResponse(w http.ResponseWriter, httpStatus int, objs interface{}) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatus)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "      ")

	if objs == nil {

		if httpStatus == http.StatusNoContent {
			return
		} else if httpStatus == http.StatusNotFound {
			if err := encoder.Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
				panic(err)
			}
		} else {
			return
		}

	} else {
		if err := encoder.Encode(objs); err != nil {
			panic(err)
		}
	}
}
