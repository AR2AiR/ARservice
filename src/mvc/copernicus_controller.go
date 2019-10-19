package mvc

import (
	"log"
	"net/http"
)

func loadCopernicusImageData(w http.ResponseWriter, r *http.Request) {
	response := queryCopernicusImageData()
	replyImageData(response, w)
}

func replyImageData(imageData []byte, w http.ResponseWriter) {
	_, err := w.Write(imageData)
	if err != nil {
		log.Print(err)
	}
}