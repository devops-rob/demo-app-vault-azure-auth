package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func metaData() (Metadata metadataJson){

	// Create HTTP requests for a managed services for Azure resources token to access Azure Resource Manager
	var metaEndpoint *url.URL
	metaEndpoint, err := url.Parse("http://169.254.169.254/metadata/instance?api-version=2021-02-01")
	if err != nil {
		fmt.Println("Error creating URL: ", err)
		return
	}
	req, err := http.NewRequest("GET", metaEndpoint.String(), nil)
	if err != nil {
		fmt.Println("Error creating HTTP request: ", err)
		return
	}
	req.Header.Add("Metadata", "true")

	// Call managed services for Azure resources token endpoint
	resp, err := Client.Do(req)
	if err != nil{
		fmt.Println("Error calling Metadata endpoint: ", err)
		return
	}

	// Pull out response body
	responseBytes,err := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body : ", err)
		return
	}

	// Unmarshall response body into struct
	err = json.Unmarshal(responseBytes, &Metadata)
	if err != nil {
		fmt.Println("Error unmarshalling the response:", err)
		return
	}

	//fmt.Printf(Metadata.Compute.SubscriptionId)
	return Metadata
}

