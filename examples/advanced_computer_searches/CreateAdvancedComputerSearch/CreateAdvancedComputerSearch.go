package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

// Define the name of the advanced computer search
const advancedComputerSearchName = "YourSearchName" // Replace with the actual name

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for Jamf Pro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new Jamf Pro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the advanced computer search details
	newSearch := &jamfpro.ResponseAdvancedComputerSearch{
		Name:   "Advanced Search Name",
		ViewAs: "Standard Web Page",
		Criteria: []jamfpro.AdvancedComputerSearchesCriteria{
			{
				Criterion: jamfpro.CriterionDetail{
					Name:         "Last Inventory Update",
					Priority:     0,
					AndOr:        "and",
					SearchType:   "more than x days ago",
					Value:        "7",
					OpeningParen: false,
					ClosingParen: false,
				},
			},
		},
		DisplayFields: []jamfpro.AdvancedComputerSearchesDisplayField{
			{
				DisplayField: jamfpro.DisplayFieldDetail{
					Name: "IP Address",
				},
			},
		},
		Site: jamfpro.AdvancedComputerSearchesSiteDetail{
			ID:   -1,
			Name: "None",
		},
	}

	// Create the advanced computer search
	createdSearch, err := client.CreateAdvancedComputerSearch(newSearch)
	if err != nil {
		fmt.Println("Error creating advanced computer search:", err)
		return
	}

	// Print the created advanced computer search details
	createdSearchXML, err := xml.MarshalIndent(createdSearch, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling created search to XML:", err)
		return
	}
	fmt.Printf("Created Advanced Computer Search:\n%s\n", string(createdSearchXML))
}
