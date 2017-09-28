package network

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func setServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprint(err)))
}


func setErrorResult(writer http.ResponseWriter, err error) {
	var result = map[string]interface{}{"error": fmt.Sprint(err)}
	json.NewEncoder(writer).Encode(&result)
}