package callback

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

func Quit() {
	req, err := http.NewRequest(http.MethodPost, "http://localhost:15000/quitquitquit", bytes.NewReader([]byte{}))
	if err != nil {
		log.Println("Invalid quit request", err)
		return
	}
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	if _, err = client.Do(req); err != nil {
		log.Println("No istio-proxy running to quit", err)
	} else {
		log.Println("Success to send quit message to istio-proxy")
	}
}
