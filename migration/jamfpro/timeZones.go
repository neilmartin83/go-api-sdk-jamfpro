// timeZones.go
// Jamf Pro Api
// API requires the structs to support JSON.

package jamfpro

import (
	"fmt"
)

const uriTimeZones = "/api/v1/time-zones"

type TimeZoneInformation struct {
	ZoneId      string `json:"zoneId"`
	Region      string `json:"region"`
	DisplayName string `json:"displayName"`
}

func (c *Client) GetTimeZoneInformation() ([]TimeZoneInformation, error) {
	var timeZones []TimeZoneInformation
	resp, err := httpclient.Get(c, uriTimeZones, &timeZones)
	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get Time Zone Information: %v", err)
	}

	return timeZones, nil
}
