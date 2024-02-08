package main

import (
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

	subscriptionID := "1" // Replace with the ID of the subscription you want to update

	// Example subscription with updated data
	updateSubscription := jamfpro.ResourceVolumePurchasingSubscription{
		Enabled:  true,
		SiteId:   "-1",
		Name:     "Example Volume Purchasing Subscription",
		Triggers: []string{"NO_MORE_LICENSES", "REMOVED_FROM_APP_STORE"},
		InternalRecipients: []jamfpro.VolumePurchasingSubscriptionSubsetInternalRecipients{
			{Frequency: "DAILY", AccountId: "1"},
		},
		ExternalRecipients: []jamfpro.VolumePurchasingSubscriptionSubsetExternalRecipients{
			{Name: "Bob Exampleman", Email: "bob@example.com"},
		},
	}

	// Call the UpdateVolumePurchasingSubscriptionByID function with the subscription ID and updated data.
	updatedSubscription, err := client.UpdateVolumePurchasingSubscriptionByID(subscriptionID, &updateSubscription)
	if err != nil {
		fmt.Printf("Error updating volume purchasing subscription with ID '%s': %s\n", subscriptionID, err)
		return
	}

	// If no error, print the details of the updated subscription.
	fmt.Printf("Updated Subscription ID: %s\n", updatedSubscription.Id)
	fmt.Printf("Name: %s\n", updatedSubscription.Name)
	// ... print other fields as needed
}
