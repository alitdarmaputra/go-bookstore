package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	json.NewDecoder(r.Body).Decode(x)
}
