package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func vaultAuth() (ClientToken VaultToken) {
	vaultEndpoint := os.Getenv("VAULT_ADDR")
	if vaultEndpoint == "" {
		log.Fatal("VAULT_ADDR environment variable not set.\n")
	}

	vaultRole := os.Getenv("VAULT_ROLE")
	if vaultEndpoint == "" {
		log.Fatal("VAULT_ROLE environment variable not set.\n")
	}

	jwt := miToken().AccessToken
	url := vaultEndpoint + "/v1/auth/azure/login"
	subscriptionId := metaData().Compute.SubscriptionId
	resourceGroup := metaData().Compute.ResourceGroupName
	vmName := metaData().Compute.Name

	payload := `
	{"role":"` + vaultRole + `",
	"jwt":"` + jwt + `", "subscription_id":"` + subscriptionId + `", "resource_group_name": "` + resourceGroup + `", "vm_name": "` + vmName + `"}`
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := Client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	err = json.Unmarshal(body, &ClientToken)
	//log.Printf("Vault token is: %s \n", ClientToken)
	return ClientToken

}

