package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func miToken() (Token tokenJson){

	// Create HTTP request for a managed services for Azure resources token to access Azure Resource Manager
	var msiEndpoint *url.URL
	msiEndpoint, err := url.Parse("http://169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01")
	if err != nil {
		fmt.Println("Error creating URL: ", err)
		return
	}
	msiParameters := msiEndpoint.Query()
	msiParameters.Add("resource", "https://management.azure.com/")
	msiEndpoint.RawQuery = msiParameters.Encode()
	req, err := http.NewRequest("GET", msiEndpoint.String(), nil)
	if err != nil {
		fmt.Println("Error creating HTTP request: ", err)
		return
	}
	req.Header.Add("Metadata", "true")

	// Call managed services for Azure resources token endpoint
	resp, err := Client.Do(req)
	if err != nil{
		fmt.Println("Error calling token endpoint: ", err)
		return
	}

	// Pull out response body
	responseBytes,err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error reading response body : ", err)
		return
	}

	// Unmarshall response body into struct
	err = json.Unmarshal(responseBytes, &Token)
	if err != nil {
		fmt.Println("Error unmarshalling the response:", err)
		return
	}

	return Token
}
