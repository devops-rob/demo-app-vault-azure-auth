package main

import "fmt"

func main() {
	token := vaultAuth().Auth.ClientToken
	fmt.Printf("Sucessfully authenticated to Vault.\nVault Token: %s \n", token)
}
