package mvc

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//type Sensors struct {
//	Id            int    `json:"id"`
//	Sampling_rate int    `json:"sampling_rate"`
//	Timestamp     string `json:"timestamp"`
//	Location struct {
//		Id int `json:"id"`
//		latitude string `json:"latitude"`
//}
//}

type Sensors struct {
	SensorData map[string]interface{}
}

func Entrypoint() {
	r := mux.NewRouter()
	r.HandleFunc("/sensors", loadAllSensorsData).Methods(http.MethodGet)
	//r.HandleFunc("/sensors/filter", getSensorsWithFilter).Methods(http.MethodGet)
	r.HandleFunc("/copernicus/image", loadCopernicusImageData).Methods(http.MethodGet);
	r.HandleFunc("/copernicus/image/v1", loadCopernicusImageDataV1).Methods(http.MethodGet);
	//r.HandleFunc("/copernicus/data", loadCopernicusData).Methods(http.MethodGet);
	log.Fatal(http.ListenAndServe(":8000", r))
}
