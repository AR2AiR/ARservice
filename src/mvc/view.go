package mvc

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func Entrypoint() {
	r := mux.NewRouter()
	r.HandleFunc("/sensors", loadAllSensorsData).Methods(http.MethodGet)
	r.HandleFunc("/sensors/filter", getSensorsWithFilter).Methods(http.MethodGet)
	r.HandleFunc("/copernicus/image", loadCopernicusImageData).Methods(http.MethodGet)
	r.HandleFunc("/copernicus/image/v1", loadCopernicusImageDataV1).Methods(http.MethodGet)
	//r.HandleFunc("/copernicus/data", loadCopernicusData).Methods(http.MethodGet);

	go func() {
		updateRateNanos := time.Second * 10
		luftDaten.updateReading()
		for {
			time.Sleep(updateRateNanos)
			luftDaten.updateReading()
		}
	}()

	log.Fatal(http.ListenAndServe(":8000", r))
}
