package mvc

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Entrypoint() {
	r := mux.NewRouter()
	r.HandleFunc("/sensors", loadAllSensorsData).Methods(http.MethodGet)
	r.HandleFunc("/sensors/filter", getSensorsWithFilter).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8000", r))
}
