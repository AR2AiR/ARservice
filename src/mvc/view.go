package mvc

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func Entrypoint() {
	r := mux.NewRouter()
	s := r.PathPrefix("/sensors/{time}").Subrouter()
	s.HandleFunc("", loadAllSensorsData).Methods(http.MethodGet)
	s.HandleFunc("/{id:[0-9]+}", getSensorById).Methods(http.MethodGet)
	s.HandleFunc("/filter", getSensorsWithFilter).Methods(http.MethodGet)

	r.HandleFunc("/copernicus/image", loadCopernicusImageData).Methods(http.MethodGet)
	r.HandleFunc("/copernicus/image/v1", loadCopernicusImageDataV1).Methods(http.MethodGet)
	//r.HandleFunc("/copernicus/data", loadCopernicusData).Methods(http.MethodGet);

	go func() {
		updateRateNanos := time.Second * 10
		luftDaten.updateReading(Last)
		luftDaten.updateReading(FiveMins)
		luftDaten.updateReading(OneHour)
		luftDaten.updateReading(TwentyFourHours)
		for {
			time.Sleep(updateRateNanos)
			luftDaten.updateReading(Last)
			luftDaten.updateReading(FiveMins)
			luftDaten.updateReading(OneHour)
			luftDaten.updateReading(TwentyFourHours)
		}
	}()

	log.Fatal(http.ListenAndServe(":8000", r))
}
