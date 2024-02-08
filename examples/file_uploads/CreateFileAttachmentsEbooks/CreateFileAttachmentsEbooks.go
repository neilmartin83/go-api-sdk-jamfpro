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

	// Define the parameters for the file upload
	resource := "ebooks" // Example resource, adjust as needed
	idType := "id"       // Can be id or name, Name is supported for all but the peripherals resource
	id := "3"            // Example ID of the resource to attach the file upload to. can be a numeral or a resource name as needed

	// Define the files to be uploaded
	files := map[string]string{
		"fileFieldName": "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/ebooks_pdf/Apple-Developer-Program-License-Agreement-20230828-English.pdf", // Replace with your actual file and field name
	}

	// Call the CreateFileAttachments method
	resp, err := client.CreateFileAttachments(resource, idType, id, files)
	if err != nil {
		fmt.Printf("Error uploading file attachments: %v\n", err)
		return
	}

	// Process the response as needed
	fmt.Println("File attachments uploaded successfully:", resp.Status)
}
