package mvc

import (
	"net/http"
	"strconv"
)

func getParamString(r *http.Request, name string, defaultValue string) string {
	queryValue := r.URL.Query().Get(name)
	if (queryValue == "") {
		return defaultValue
	}
	return queryValue
}

func getParamFloat64(r *http.Request, name string, defaultValue float64) float64 {
	queryValue := r.URL.Query().Get(name)
	if (queryValue == "") {
		return defaultValue
	} 

	value, _ := strconv.ParseFloat(queryValue, 64)
	return value
}

func getParamInt(r *http.Request, name string, defaultValue int64) int64 {
	queryValue := r.URL.Query().Get(name)
	if (queryValue == "") {
		return defaultValue
	} 
	value, _ := strconv.ParseInt(queryValue, 10, 64)
	return value
}