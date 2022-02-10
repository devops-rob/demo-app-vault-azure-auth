package main

type tokenJson struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn string `json:"expires_in"`
	ExpiresOn string `json:"expires_on"`
	NotBefore string `json:"not_before"`
	Resource string `json:"resource"`
	TokenType string `json:"token_type"`
}

type metadataJson struct {
	Compute computeJson `json:"compute"`
}

type computeJson struct {
	SubscriptionId string `json:"subscriptionId"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resourceGroupName"`
}

type VaultToken struct {
	Auth        Auth `json:"auth"`
}

type Auth struct {
	ClientToken string `json:"client_token"`
}


