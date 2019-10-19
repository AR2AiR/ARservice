package mvc

import (
	"net/http"
	"strconv"
)

func getParamFloat64(r *http.Request, name string) float64 {
	value, _ := strconv.ParseFloat(r.URL.Query().Get(name), 64)
	return value
}

func getParamInt(r *http.Request, name string) int64 {
	value, _ := strconv.ParseInt(r.URL.Query().Get(name), 2, 64)
	return value
}