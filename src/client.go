package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/**
 * Predefine
 **/
const (
	ApiURL   = "http://ip-api.com/json/"
	ApiAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:27.0) Gecko/20100101 Firefox/27.0"
)

/**
 * QueryResponse Json Object
 **/
type QueryReponse struct {
	AS          string `json:"as"`
	City        string `json:"city"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	ISP         string `json:"isp"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	ORG         string `json:"org"`
	IP          string `json:"query"`
	Region      string `json:"region"`
	RegionName  string `json:"regionName"`
        Status      string `json:"status"`
	Timezone    string `json:"timezone"`
	Zip         string `json:"zip"`
}

/**
 *  Get HTTP Response Header
 **/
func GetHeaders() http.Header {
	headers := http.Header{}
	headers.Set("Content-type", "application/json")
	return headers
}

/**
 * Get Json Object
 **/
func GetQuery(IP string) QueryReponse {

	// create a empty object
	result := QueryReponse{}

	// if empty, return
	if len(IP) == 0 {
		return result
	}

	// create a client object
	client := &http.Client{}

	// create a request object
	req, err := http.NewRequest("GET", ApiURL+IP, nil)
	if err != nil {
		return result
	}

	// apply user-agent
	req.Header.Add("User-Agent", ApiAgent)

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return result
	}

	// prepare buffer
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result
	}

	// close request
	resp.Body.Close()

	// json decode
	json.Unmarshal([]byte(buf), &result)
	return result
}
