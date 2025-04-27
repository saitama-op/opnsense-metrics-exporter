package api

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func FetchOPNSenseData(apiUrl, apiKey, apiSecret string, ignore_ssl bool) (map[string]interface{}, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: ignore_ssl, // WARNING: disables cert validation
			},
		},
	}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(apiKey, apiSecret)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	log.Println("Response Status:", resp.Status)
	log.Println("Response Headers:", resp.Header)
	if resp.StatusCode != http.StatusOK {
		log.Println("Error: Received non-200 response code")
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println("Error decoding JSON response:", err)
		return nil, err
	}

	return response, nil
}
