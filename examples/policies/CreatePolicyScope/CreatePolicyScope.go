package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define a new policy with all required fields
	newPolicy := &jamfpro.ResourcePolicy{
		General: jamfpro.PolicySubsetGeneral{
			Name:                       "jamfpro-sdk-scope-policy-config",
			Enabled:                    false,
			Trigger:                    "EVENT",
			TriggerCheckin:             false,
			TriggerEnrollmentComplete:  false,
			TriggerLogin:               false,
			TriggerLogout:              false,
			TriggerNetworkStateChanged: false,
			TriggerStartup:             false,
			Frequency:                  "Once per computer",
			RetryEvent:                 "none",
			RetryAttempts:              -1,
			NotifyOnEachFailedRetry:    false,
			LocationUserOnly:           false,
			TargetDrive:                "/",
			Offline:                    false,
			Category: jamfpro.PolicyCategory{
				ID:        -1,
				Name:      "No category assigned",
				DisplayIn: false,
				FeatureIn: false,
			},
		},
		Scope: jamfpro.PolicySubsetScope{
			AllComputers: false,
			AllJSSUsers:  false,
			Computers:    []jamfpro.PolicyDataSubsetComputer{}, // Empty as per XML
			ComputerGroups: []jamfpro.PolicyDataSubsetComputerGroup{
				{
					ID:   26,
					Name: "Test-Smart-Group-1",
				},
				// Additional computer groups can be added here if needed
			},
			JSSUsers:      []jamfpro.PolicyDataSubsetJSSUser{},      // Empty as per XML
			JSSUserGroups: []jamfpro.PolicyDataSubsetJSSUserGroup{}, // Empty as per XML
			Buildings: []jamfpro.PolicyDataSubsetBuilding{
				{
					ID:   1320,
					Name: "Apple Park 2",
				},
				// Additional buildings can be added here if needed
			},
			Departments: []jamfpro.PolicyDataSubsetDepartment{
				{
					ID:   23515,
					Name: "JLtestDept2",
				},
				// Additional departments can be added here if needed
			},
			LimitToUsers: jamfpro.PolicyLimitToUsers{
				UserGroups: []string{}, // Empty as per XML
				// Additional user groups can be added here if needed
			},
			Limitations: jamfpro.PolicySubsetScopeLimitations{
				Users:           []jamfpro.PolicyDataSubsetUser{},           // Empty as per XML
				UserGroups:      []jamfpro.PolicyDataSubsetUserGroup{},      // Empty as per XML
				NetworkSegments: []jamfpro.PolicyDataSubsetNetworkSegment{}, // Empty as per XML
				IBeacons:        []jamfpro.PolicyDataSubsetIBeacon{},        // Empty as per XML
			},
			Exclusions: jamfpro.PolicySubsetScopeExclusions{
				Computers:       []jamfpro.PolicyDataSubsetComputer{},       // Empty as per XML
				ComputerGroups:  []jamfpro.PolicyDataSubsetComputerGroup{},  // Empty as per XML
				Users:           []jamfpro.PolicyDataSubsetUser{},           // Empty as per XML
				UserGroups:      []jamfpro.PolicyDataSubsetUserGroup{},      // Empty as per XML
				Buildings:       []jamfpro.PolicyDataSubsetBuilding{},       // Empty as per XML
				Departments:     []jamfpro.PolicyDataSubsetDepartment{},     // Empty as per XML
				NetworkSegments: []jamfpro.PolicyDataSubsetNetworkSegment{}, // Empty as per XML
				JSSUsers:        []jamfpro.PolicyDataSubsetJSSUser{},        // Empty as per XML
				JSSUserGroups:   []jamfpro.PolicyDataSubsetJSSUserGroup{},   // Empty as per XML
				IBeacons:        []jamfpro.PolicyDataSubsetIBeacon{},        // Empty as per XML
			},
		},
		SelfService: jamfpro.PolicySubsetSelfService{
			UseForSelfService:           false,
			SelfServiceDisplayName:      "",
			InstallButtonText:           "Install",
			ReinstallButtonText:         "",
			SelfServiceDescription:      "",
			ForceUsersToViewDescription: false,
			//SelfServiceIcon:             jamfpro.Icon{ID: -1, Filename: "", URI: ""},
			FeatureOnMainPage: false,
		},
		AccountMaintenance: jamfpro.PolicySubsetAccountMaintenance{
			ManagementAccount: jamfpro.PolicySubsetAccountMaintenanceManagementAccount{
				Action:                "doNotChange",
				ManagedPassword:       "",
				ManagedPasswordLength: 0,
			},
			OpenFirmwareEfiPassword: jamfpro.PolicySubsetAccountMaintenanceOpenFirmwareEfiPassword{
				OfMode:           "none",
				OfPassword:       "",
				OfPasswordSHA256: "",
			},
		},
		Maintenance: jamfpro.PolicySubsetMaintenance{
			Recon:                    false,
			ResetName:                false,
			InstallAllCachedPackages: false,
			Heal:                     false,
			Prebindings:              false,
			Permissions:              false,
			Byhost:                   false,
			SystemCache:              false,
			UserCache:                false,
			Verify:                   false,
		},
		FilesProcesses: jamfpro.PolicySubsetFilesProcesses{
			DeleteFile:           false,
			UpdateLocateDatabase: false,
			SpotlightSearch:      "",
			SearchForProcess:     "",
			KillProcess:          false,
			RunCommand:           "",
		},
		// User interation policy settings
		UserInteraction: jamfpro.PolicySubsetUserInteraction{
			MessageStart:          "",
			AllowUserToDefer:      true,
			AllowDeferralUntilUtc: "",
			AllowDeferralMinutes:  0,
			MessageFinish:         "",
		},
		Reboot: jamfpro.PolicySubsetReboot{
			Message:                     "This computer will restart in 5 minutes. Please save anything you are working on and log out by choosing Log Out from the bottom of the Apple menu.",
			StartupDisk:                 "Current Startup Disk",
			SpecifyStartup:              "",
			NoUserLoggedIn:              "Do not restart",
			UserLoggedIn:                "Do not restart",
			MinutesUntilReboot:          5,
			StartRebootTimerImmediately: false,
			FileVault2Reboot:            false,
		},
	}

	policyXML, err := xml.MarshalIndent(newPolicy, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling policy data: %v", err)
	}
	fmt.Println("Policy Details to be Sent:\n", string(policyXML))

	// Call CreatePolicy function
	createdPolicy, err := client.CreatePolicy(newPolicy)
	if err != nil {
		log.Fatalf("Error creating policy: %v", err)
	}

	// Pretty print the created policy details in XML
	policyXML, err = xml.MarshalIndent(createdPolicy, "", "    ") // Indent with 4 spaces and use '='
	if err != nil {
		log.Fatalf("Error marshaling policy details data: %v", err)
	}
	fmt.Println("Created Policy Details:\n", string(policyXML))
}
