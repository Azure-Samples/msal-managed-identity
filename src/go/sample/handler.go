package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	mi "github.com/AzureAD/microsoft-authentication-library-for-go/apps/managedidentity"
)

func getSecretFromAzureVault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "going to try get secret from managed identity")

	keyVaultUri := "your-key-vault-uri"
	secretName := "your-secret-name"

	// Uncomment one of the following lines to create a new managed identity client for UserAssigned. Will need to create UserAssigned managed identity on Azure
	miClient, err := mi.New(mi.SystemAssigned())
	// miClient, err := mi.New(mi.UserAssignedClientID("my-client-id"))
	// miClient, err := mi.New(mi.UserAssignedObjectID("my-object-id"))
	// miClient, err := mi.New(mi.UserAssignedResourceID("my-resource-id"))
	if err != nil {
		fmt.Fprintf(w, "failed to create a new managed identity client: %v", err)
		return
	}

	accessToken, err := miClient.AcquireToken(context.Background(), "https://vault.azure.net")
	if err != nil {
		fmt.Fprintf(w, "failed to acquire token: %v", err)
		return
	}

	fmt.Fprintf(w, "Access token: %s", accessToken.AccessToken)

	// Create http request using access token
	url := fmt.Sprintf("%ssecrets/%s?api-version=7.2", keyVaultUri, secretName)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(w, "Error creating request: %v", err)
	}

	// Set the authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken.AccessToken))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading response body: %v", err)
	}

	// Combine all received buffer streams into one buffer, and then into a string
	var parsedData map[string]interface{}
	if err := json.Unmarshal(body, &parsedData); err != nil {
		fmt.Fprintf(w, "Error parsing JSON: %v", err)
	}

	// Print the response body
	fmt.Fprintf(w, "The secret, %s, has a value of: %s", secretName, string(body))
}

func main() {
	http.HandleFunc("/api/AcquireTokenMSI", getSecretFromAzureVault)
	http.ListenAndServe(":8080", nil)
}
