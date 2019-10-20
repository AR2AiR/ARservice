package mvc

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func loadCopernicusImageForecast(w http.ResponseWriter, r *http.Request) {

	service := getParamString(r, "service", "ENSEMBLE-FORECAST")
	species := getParamString(r, "species", "PM10")
	startLatitude := getParamFloat64(r, "startLatitude", 30.05)
	endLatitude := getParamFloat64(r, "endLatitude", -24.95)
	startLongitude := getParamFloat64(r, "startLongitude", 71.95)
	endLongitude := getParamFloat64(r, "endLongitude", 44.95)
	level := getParamInt(r, "level", 0)
	width := getParamInt(r, "width", 810)
	height := getParamInt(r, "height", 495)

	response := queryCopernicusImageData(service, species, startLatitude, endLatitude, startLongitude, endLongitude, level, width, height)
	replyImageData(response, w)
}

func loadCopernicusImageDataV1(w http.ResponseWriter, r *http.Request) {
	response := queryCopernicusImageDataV1()
	replyImageData(response, w)
}
func replyImageData(imageData []byte, w http.ResponseWriter) {
	_, err := w.Write(imageData)
	if err != nil {
		log.Print(err)
	}
}

func queryCopernicusImageData(service string, species string, startLatitude float64, endLatitude float64, startLongitude float64, endLongitude float64, level int64, width int64, height int64) []byte {

	const apiToken = "__-Tlb4srqXM2EzZkgj_Va5OdozWR9G-dW7TS0yFsZL0I__"
	serviceURL := fmt.Sprintf("https://geoservices.regional.atmosphere.copernicus.eu/services/CAMS50-%s-01-EUROPE-WMS", service)

	baseURL, err := url.Parse(serviceURL)
	if err != nil {
		log.Fatal("Malformed URL: ", err.Error())
	}

	params := url.Values {}
	params.Add("service", "WMS")
	params.Add("VERSION", "1.3.0")
	params.Add("LAYERS", fmt.Sprintf("%s__SPECIFIC_HEIGHT_LEVEL_ABOVE_GROUND", species))
	params.Add("STYLES", fmt.Sprintf("%s_USI__HEIGHT__SHADING", species))
	params.Add("DIM_REFERENCE_TIME", "2019-10-17T00:00:00Z")
	params.Add("TIME", "2019-10-17T05:00:00Z")
	params.Add("ELEVATION", fmt.Sprintf("%d", level))
	params.Add("FORMAT", "image/png")
	params.Add("TOKEN", apiToken)
	params.Add("REQUEST", "GetMap")
	params.Add("CRS", "EPSG:4326")
	params.Add("BBOX", fmt.Sprintf("%.2f,%.2f,%.2f,%.2f", startLatitude, endLatitude, startLongitude, endLongitude))
	params.Add("WIDTH", fmt.Sprintf("%d", width))
	params.Add("HEIGHT", fmt.Sprintf("%d", height))

	url, err := url.QueryUnescape(params.Encode())
	if err != nil {
		log.Fatal(err.Error())
	}
	baseURL.RawQuery = url

	// log.Printf("query %s", url)
	log.Printf("url %s", baseURL.String())

	response, err := http.Get(baseURL.String())
	if err != nil {
		log.Printf("Error fetching copernicus image data: %s\n", err)
	}

	imageData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response image %s\n", err)
	}

	return imageData
}

func queryCopernicusImageDataV1() []byte {

	const apiToken = "__-Tlb4srqXM2EzZkgj_Va5OdozWR9G-dW7TS0yFsZL0I__"
	const serviceURL = "https://geoservices.regional.atmosphere.copernicus.eu/services/CAMS50-ENSEMBLE-FORECAST-01-EUROPE-WMS?&service=WMS&VERSION=1.3.0&LAYERS=PM10__SPECIFIC_HEIGHT_LEVEL_ABOVE_GROUND&STYLES=PM10_USI__HEIGHT__SHADING&DIM_REFERENCE_TIME=2019-10-17T00:00:00Z&TIME=2019-10-17T05:00:00Z&ELEVATION=0&FORMAT=image/png&TOKEN=__-Tlb4srqXM2EzZkgj_Va5OdozWR9G-dW7TS0yFsZL0I__&REQUEST=GetMap&CRS=EPSG:4326&BBOX=30.05,-24.95,71.95,44.95&WIDTH=810&HEIGHT=495";

	response, err := http.Get(serviceURL)
	if err != nil {
		log.Printf("Error fetching copernicus image data: %s\n", err)
	}

	imageData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response image %s\n", err)
	}

	return imageData
}