package mvc

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var luftDaten LuftDatenReading

type LuftDatenReading struct {
	lastReading *Welcome
	mut sync.Mutex
}

func (c *LuftDatenReading)queryAllSensorsData() *Welcome {
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

func (c *LuftDatenReading) updateReading() {
	c.mut.Lock()
	c.lastReading = c.queryAllSensorsData()
	c.mut.Unlock()
	log.Println("Updated Luftdaten data!")
}

func (c *LuftDatenReading) getLastReading() *Welcome {
	c.mut.Lock()
	data := c.lastReading
	c.mut.Unlock()
	return data
}


