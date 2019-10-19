package mvc

import (
	"io/ioutil"
	"log"
	"net/http"
)

func queryCopernicusImageData() []byte {

	const apiToken = "__-Tlb4srqXM2EzZkgj_Va5OdozWR9G-dW7TS0yFsZL0I__"
	const serviceURL = "https://geoservices.regional.atmosphere.copernicus.eu/services/CAMS50-ENSEMBLE-FORECAST-01-EUROPE-WMS?&service=WMS&VERSION=1.3.0&LAYERS=PM10__SPECIFIC_HEIGHT_LEVEL_ABOVE_GROUND&STYLES=PM10_USI__HEIGHT__SHADING&DIM_REFERENCE_TIME=2019-10-15T00:00:00Z&TIME=2019-10-15T05:00:00Z&ELEVATION=0&FORMAT=image/png&TOKEN=__-Tlb4srqXM2EzZkgj_Va5OdozWR9G-dW7TS0yFsZL0I__&REQUEST=GetMap&CRS=EPSG:4326&BBOX=30.05,-24.95,71.95,44.95&WIDTH=810&HEIGHT=495"

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