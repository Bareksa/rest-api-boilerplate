package utils

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, httpStatus int, data interface{})  {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(data)
}
