package main

import (
	"encoding/xml"
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
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

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

	// Define the advanced computer search details
	updatedSearch := &jamfpro.ResourceAdvancedComputerSearch{
		Name:   "jamf api sdk advanced search",
		ViewAs: "Standard Web Page",
		Criteria: jamfpro.SharedContainerCriteria{
			Size: 1,
			Criterion: []jamfpro.SharedSubsetCriteria{
				{
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
		DisplayFields: []jamfpro.SharedAdvancedSearchContainerDisplayField{
			{
				DisplayField: []jamfpro.SharedAdvancedSearchSubsetDisplayField{
					{
						Name: "Activation Lock Manageable",
					},
					{
						Name: "Apple Silicon",
					},
					{
						Name: "Architecture Type",
					},
					{
						Name: "Available RAM Slots",
					},
				},
			},
		},
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
	}

	searchID := 12

	// Convert the profile to XML to see the output (optional, for debug purposes)
	xmlData, err := xml.MarshalIndent(updatedSearch, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling XML: %v", err)
	}
	fmt.Printf("XML Request: %s\n", xmlData)

	// Create the advanced computer search
	updatedSearchResp, err := client.UpdateAdvancedComputerSearchByID(searchID, updatedSearch)
	if err != nil {
		fmt.Println("Error creating advanced computer search:", err)
		return
	}
	// Print the created advanced computer search details
	createdSearchXML, err := xml.MarshalIndent(updatedSearchResp, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling created search to XML:", err)
		return
	}
	fmt.Printf("Created Advanced Computer Search:\n%s\n", string(createdSearchXML))
}
