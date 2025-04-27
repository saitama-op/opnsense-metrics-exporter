package api

import (
    "encoding/json"
    "net/http"
)

func FetchOPNSenseData(apiUrl, apiKey string) (map[string]interface{}, error) {
    req, err := http.NewRequest("GET", apiUrl, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Add("Authorization", "Bearer " + apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var response map[string]interface{}
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        return nil, err
    }

    return response, nil
}
