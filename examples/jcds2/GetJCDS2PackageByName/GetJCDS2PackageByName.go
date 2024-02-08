package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := logger.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Let's assume you want to get the uploaded package with name "Firefox 122.0.dmg"
	packageName := "2"

	configuration, err := client.GetJCDS2PackageByName(packageName)
	if err != nil {
		log.Fatalf("Error fetching JCDS 2.0 package by name: %v", err)
	}

	// Print the configuration in a pretty JSON format
	configJSON, err := json.MarshalIndent(configuration, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling configuration data: %v", err)
	}
	fmt.Printf("Fetched JCDS 2.0 package by Name:\n%s\n", configJSON)
}
