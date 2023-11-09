// classicapi_classes.go
// Jamf Pro Classic Api - Classes
// api reference: https://developer.jamf.com/jamf-pro/reference/classes
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

// Constants for the classes endpoint
const uriClasses = "/JSSResource/classes"

// ResponseClassesList represents the XML response for a list of classes.
type ResponseClassesList struct {
	Size    int         `xml:"size"`
	Classes []ClassItem `xml:"class"`
}

// ClassItem represents a single class item in the list.
type ClassItem struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	Description string `xml:"description"`
}

// Structs for the Class response by ID
type ResponseClasses struct {
	ID                  int                 `xml:"id,omitempty"`
	Source              string              `xml:"source,omitempty"`
	Name                string              `xml:"name"` // Required
	Description         string              `xml:"description,omitempty"`
	Site                ClassSite           `xml:"site"`
	MobileDeviceGroup   ClassDeviceGroup    `xml:"mobile_device_group,omitempty"`
	Students            []ClassStudent      `xml:"students>student,omitempty"`
	Teachers            []ClassTeacher      `xml:"teachers>teacher,omitempty"`
	TeacherIDs          []ClassTeacherID    `xml:"teacher_ids>id,omitempty"`
	StudentGroupIDs     []ClassGroupID      `xml:"student_group_ids>id,omitempty"`
	TeacherGroupIDs     []ClassGroupID      `xml:"teacher_group_ids>id,omitempty"`
	MobileDevices       []ClassMobileDevice `xml:"mobile_devices>mobile_device,omitempty"`
	MobileDeviceGroupID []ClassGroupID      `xml:"mobile_device_group_id>id,omitempty"`
	MeetingTimes        ClassMeetingTimes   `xml:"meeting_times,omitempty"`
	AppleTVs            []ClassAppleTV      `xml:"apple_tvs>apple_tv,omitempty"`
}
type ClassSite struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name"` // Required
}

type ClassDeviceGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type ClassStudent struct {
	Student string `xml:"student,omitempty"`
}

type ClassTeacher struct {
	Teacher string `xml:"teacher,omitempty"`
}

type ClassTeacherID struct {
	ID int `xml:"id,omitempty"`
}

type ClassGroupID struct {
	ID int `xml:"id,omitempty"`
}

type ClassMobileDevice struct {
	Name           string `xml:"name,omitempty"`
	UDID           string `xml:"udid,omitempty"`
	WifiMacAddress string `xml:"wifi_mac_address,omitempty"`
}

type ClassMeetingTimes struct {
	MeetingTime ClassMeetingTime `xml:"meeting_time,omitempty"`
}

type ClassMeetingTime struct {
	Days      string `xml:"days,omitempty"`
	StartTime int    `xml:"start_time,omitempty"`
	EndTime   int    `xml:"end_time,omitempty"`
}

type ClassAppleTV struct {
	Name            string `xml:"name,omitempty"`
	UDID            string `xml:"udid,omitempty"`
	WifiMacAddress  string `xml:"wifi_mac_address,omitempty"`
	DeviceID        string `xml:"device_id,omitempty"`
	AirplayPassword string `xml:"airplay_password,omitempty"`
}

// GetClasses gets a list of all classes.
func (c *Client) GetClasses() (*ResponseClassesList, error) {
	endpoint := uriClasses

	var classes ResponseClassesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &classes)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Classes: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &classes, nil
}

// GetClassesByID retrieves a class by its ID.
func (c *Client) GetClassesByID(id int) (*ResponseClasses, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriClasses, id)

	var class ResponseClasses
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &class)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Class by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &class, nil
}

// GetClassesByName retrieves a class by its name.
func (c *Client) GetClassesByName(name string) (*ResponseClasses, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriClasses, name)

	var class ResponseClasses
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &class)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Class by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &class, nil
}

// CreateClassesByID creates a new class with the given details.
func (c *Client) CreateClassesByID(class *ResponseClasses) (*ResponseClasses, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriClasses) // Using ID 0 for creation as per API pattern

	// If the site is not provided, set default values
	if class.Site.ID == 0 && class.Site.Name == "" {
		class.Site = ClassSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Wrap the class request with the desired XML structure using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"class"`
		*ResponseClasses
	}{
		ResponseClasses: class,
	}

	var createdClass ResponseClasses
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdClass)
	if err != nil {
		return nil, fmt.Errorf("failed to create Class: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdClass, nil
}

// UpdateClassByID updates an existing class with the given ID.
func (c *Client) UpdateClassesByID(id int, class *ResponseClasses) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriClasses, id)

	// Wrap the class request with the desired XML structure using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"class"`
		*ResponseClasses
	}{
		ResponseClasses: class,
	}

	_, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, nil)
	if err != nil {
		return fmt.Errorf("failed to update Class by ID: %v", err)
	}

	return nil
}

// UpdateClassByName updates an existing class with the given name.
func (c *Client) UpdateClassesByName(name string, class *ResponseClasses) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriClasses, name)

	// Wrap the class request with the desired XML structure using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"class"`
		*ResponseClasses
	}{
		ResponseClasses: class,
	}

	_, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, nil)
	if err != nil {
		return fmt.Errorf("failed to update Class by Name: %v", err)
	}

	return nil
}

// DeleteClassByID deletes an existing class with the given ID.
func (c *Client) DeleteClassByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriClasses, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Class by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteClassByName deletes a class by its name.
func (c *Client) DeleteClassByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriClasses, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Class by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
