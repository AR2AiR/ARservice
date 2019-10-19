package mvc

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func loadAllSensorsData(w http.ResponseWriter, r *http.Request) {
	response := queryAllSensorsData()
	reply(response, w)
}

//func getSensorsWithFilter(w http.ResponseWriter, r *http.Request) {
//	param := r.URL.Query().Get("area")
//	lat, lon, dist := processAreaQueryStringParams(param)
//	response := queryAllSensorsData()
//	var sensorsInArea []SensorReading
//	for _, sensor := range response {
//		sensorLat := string2Float64(sensor.Location.Latitude)
//		sensorLon := string2Float64(sensor.Location.Longitude)
//		if distanceMetersBetweenTwoLocations(lat, lon, sensorLat, sensorLon) < dist {
//			sensorsInArea = append(sensorsInArea, sensor)
//		}
//	}
//	fmt.Println(lat, lon, dist)
//	fmt.Println(response)
	//reply(sensorsInArea, w)
//}

func processAreaQueryStringParams(param string) (float64, float64, float64) {
	splitParams := strings.Split(param, ",")
	lat, _ := strconv.ParseFloat(splitParams[0], 64)
	lon, _ := strconv.ParseFloat(splitParams[1], 64)
	dist, _ := strconv.ParseFloat(splitParams[2], 64)
	return lat, lon, dist
}

func distanceMetersBetweenTwoLocations(lat1, lon1, lat2, lon2 float64) float64 {
	earthRadiusKms := 6378.137
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0
	a := math.Sin(dLat/2.0)*math.Sin(dLat/2.0) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := earthRadiusKms * c
	return d * 1000
}

func string2Float64(numStr string) float64 {
	num, _ := strconv.ParseFloat(numStr, 64)
	return num
}

func reply(sensorData *Welcome, w http.ResponseWriter) {
	bytes, err := sensorData.Marshal()
	if err != nil {
		log.Print(err)
	}
	_, err = w.Write(bytes)
	if err != nil {
		log.Print(err)
	}
}
