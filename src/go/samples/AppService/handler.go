package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	mi "github.com/AzureAD/microsoft-authentication-library-for-go/apps/managedidentity"
)

func getSecretFromAzureVault() string {
	keyVaultUri := "https://go-lang-kv.vault.azure.net/"
	secretName := "go-lang-secret"
	finalResult := ""

	miClient, err := mi.New(mi.SystemAssigned())

	if err != nil {
		log.Printf("failed to create a new managed identity client: %v", err)
	}

	source, err := mi.GetSource()

	if err != nil {
		log.Printf("failed to get source: %v", err)
	}
	fmt.Println("Managed Identity Source: ", source)

	accessTokenR, err := miClient.AcquireToken(context.Background(), "https://vault.azure.net")
	if err != nil {
		finalResult += "\n::error result for resource id::" + err.Error()
		log.Printf("failed to acquire token: %v", err)
	}
	finalResult += "\n::got result for resource id::" + accessTokenR.ExpiresOn.String()

	// Create http request using access token
	url := fmt.Sprintf("%ssecrets/%s?api-version=7.2", keyVaultUri, secretName)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
	}

	// Set the authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessTokenR.AccessToken))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
	}

	// Combine all received buffer streams into one buffer, and then into a string
	var parsedData map[string]interface{}
	if err := json.Unmarshal(body, &parsedData); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Print the response body
	return fmt.Sprintf(":: %s :: The secret from Object , %s, has a value of: %s", finalResult, secretName, parsedData["value"])
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "response : is : "+getSecretFromAzureVault())
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/mifunction", helloHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
