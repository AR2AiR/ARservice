package mvc

import (
	"io/ioutil"
	"log"
	"net/http"
)

type SensorDataValues struct {
	Id        int    `json:"id"`
	ValueType string `json:"value_type"`
	Value     float64 `json:",string"`
}

type SensorReading struct {
	DataValue    []SensorDataValues `json:"sensordatavalues"`
	Timestamp    string `json:"timestamp"`
	Id           int    `json:id`
	SamplingRate int    `json:"sampling_rate"`
	Location     struct {
		Altitude      string `json:altitude`
		Id            int    `json:"id"`
		Latitude      string `json:"latitude"`
		Longitude     string `json:"longitude"`
		ExactLocation int    `json:"exact_location"`
		Country       string `json:"country"`
		Indoor        int    `json:"indoor"`
	}
	Sensor struct {
		Pin        string `json:"pin"`
		Id         int    `json:"id"`
		SensorType struct {
			Id           int    `json:"id"`
			Name         string `json:"name"`
			Manufacturer string `json:"manufacturer"`
		}
	}
}

func queryAllSensorsData() *Welcome {
	response, err := http.Get("https://api.luftdaten.info/static/v1/data.json")
	if err != nil {
		log.Printf("Error fetching sensor data %s\n", err)
	}

	var sensorData Welcome
	bodyBytes, err := ioutil.ReadAll(response.Body)
	sensorData, err = UnmarshalWelcome(bodyBytes)
	if err != nil {
		log.Printf("Error converting response to json %s\n", err)
	}
	return &sensorData
}
