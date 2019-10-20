package mvc

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var luftDaten LuftDatenReading

const (
	Last            string = "last"
	FiveMins        string = "5mins"
	OneHour         string = "1hour"
	TwentyFourHours string = "24hours"
)

type LuftDatenReading struct {
	lastReading    *Welcome
	mutLastReading sync.Mutex

	avgLast5MinsReading *Welcome
	mutAvgLast5Mins     sync.Mutex

	avgLast1HourReading *Welcome
	mutAvgLast1hour     sync.Mutex

	avgLast24HoursReading *Welcome
	mutAvgLast24Hours     sync.Mutex
}

func (c *LuftDatenReading) queryAllSensorsData(uri string) *Welcome {
	response, err := http.Get(uri)
	if err != nil {
		log.Printf("Error fetching sensor data %s\n", err)
	}

	var sensorData Welcome
	bodyBytes, err := ioutil.ReadAll(response.Body)
	sensorData, err = UnmarshalWelcome(bodyBytes)
	if err != nil {
		log.Printf("Error converting response to json %s\n", err)
	}
	trimedData := postProcessSensorsData(&sensorData)
	return trimedData
}

func (c *LuftDatenReading) updateReading(tp string) {

	switch tp {
	case Last:
		c.mutLastReading.Lock()
		c.lastReading = c.queryAllSensorsData("https://api.luftdaten.info/static/v1/data.json")
		c.mutLastReading.Unlock()
		break
	case FiveMins:
		c.mutAvgLast5Mins.Lock()
		c.avgLast5MinsReading = c.queryAllSensorsData("http://api.luftdaten.info/static/v2/data.json")
		c.mutAvgLast5Mins.Unlock()
		break
	case OneHour:
		c.mutAvgLast1hour.Lock()
		c.avgLast1HourReading = c.queryAllSensorsData("http://api.luftdaten.info/static/v2/data.1h.json")
		c.mutAvgLast1hour.Unlock()
		break
	case TwentyFourHours:
		c.mutAvgLast24Hours.Lock()
		c.avgLast24HoursReading = c.queryAllSensorsData("http://api.luftdaten.info/static/v2/data.24h.json")
		c.mutAvgLast24Hours.Unlock()
		break
	}
	log.Printf("Updated luftdaten data type: %s", tp)

}

func (c *LuftDatenReading) getLastReading(tp string) *Welcome {

	var data *Welcome
	switch tp {
	case Last:
		c.mutLastReading.Lock()
		data = c.lastReading
		c.mutLastReading.Unlock()
		break
	case FiveMins:
		c.mutAvgLast5Mins.Lock()
		data = c.avgLast5MinsReading
		c.mutAvgLast5Mins.Unlock()
		break
	case OneHour:
		c.mutAvgLast1hour.Lock()
		data = c.avgLast1HourReading
		c.mutAvgLast1hour.Unlock()
		break
	case TwentyFourHours:
		c.mutAvgLast24Hours.Lock()
		data = c.avgLast24HoursReading
		c.mutAvgLast24Hours.Unlock()
		break
	default:
		log.Print("Wrong type of sensor reading requested")
	}
	return data
}

func postProcessSensorsData(data *Welcome) *Welcome {
	var trimedData Welcome
	for _, element := range *data {
		if sensorHasPmReading(&element) {
			trimedData = append(trimedData, element)
		}
	}
	return &trimedData
}

func sensorHasPmReading(elem *WelcomeElement) bool {
	sensorName := elem.Sensor.SensorType.Name
	if sensorName == Sds011 || sensorName == Sds021 || sensorName == Hpm ||
		sensorName == Pms1003 || sensorName == Pms3003 || sensorName == Pms5003 ||
		sensorName == Pms7003 || sensorName == Ppd42NS {
		return true
	}
	return false
}

//func getHTMLpage() {
//	file, _ := http.Get("http://archive.luftdaten.info/2019-10-18/")
//	fmt.Print(file.Body)
//}
