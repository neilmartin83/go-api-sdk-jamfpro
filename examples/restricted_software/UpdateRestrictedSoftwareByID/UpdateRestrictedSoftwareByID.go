package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	//logLevel := logger.LogLevelInfo // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	restrictedSoftwareID := 1 // Replace with actual ID

	updatedRestrictedSoftware := &jamfpro.ResourceRestrictedSoftware{
		General: jamfpro.RestrictedSoftwareSubsetGeneral{
			Name:                  "Restrict High Sierra",
			ProcessName:           "Install macOS High Sierra.app",
			MatchExactProcessName: true,
			SendNotification:      true,
			KillProcess:           true,
			DeleteExecutable:      true,
			DisplayMessage:        "High Sierra is not yet supported, check Self Service after public release.",
			Site: jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: jamfpro.RestrictedSoftwareSubsetScope{
			AllComputers:   false,
			Computers:      []jamfpro.RestrictedSoftwareSubsetScopeComputer{},
			ComputerGroups: []jamfpro.RestrictedSoftwareSubsetScopeComputerGroup{},
			Buildings:      []jamfpro.RestrictedSoftwareSubsetScopeBuilding{},
			Departments:    []jamfpro.RestrictedSoftwareSubsetScopeDepartment{},
			Exclusions: jamfpro.RestrictedSoftwareSubsetScopeExclusions{
				Computers:      []jamfpro.RestrictedSoftwareSubsetScopeComputer{},
				ComputerGroups: []jamfpro.RestrictedSoftwareSubsetScopeComputerGroup{},
				Buildings:      []jamfpro.RestrictedSoftwareSubsetScopeBuilding{},
				Departments:    []jamfpro.RestrictedSoftwareSubsetScopeDepartment{},
				Users:          []jamfpro.RestrictedSoftwareSubsetScopeUser{},
			},
		},
	}

	err = client.UpdateRestrictedSoftwareByID(restrictedSoftwareID, updatedRestrictedSoftware)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Restricted software updated successfully.")
}
