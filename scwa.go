package scwa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseURL = "https://cp-%s.scaleway.com/products/servers/availability"

type getAvailabilitiesResponse struct {
	Servers map[string]struct {
		Availability string `json:"availability"`
	} `json:"servers"`
}

// GetAvailabilities returns availabilities by region
func GetAvailabilities(region string) (availabilities map[string]string, err error) {
	response, err := http.Get(fmt.Sprintf(baseURL, region))
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var r getAvailabilitiesResponse

	err = json.Unmarshal(body, &r)
	if err != nil {
		return
	}

	availabilities = make(map[string]string, len(r.Servers))

	for k, v := range r.Servers {
		availabilities[k] = v.Availability
	}
	return
}

// GetAvailability returns server availability by flavor (eg GP1-XS, START1-L, ....) and région (par1 or asm1)
func GetAvailability(region, flavor string) (availability string, err error) {
	flavor = strings.ToUpper(flavor)

	availabilties, err := GetAvailabilities(region)
	if err != nil {
		return "", err
	}

	availability = availabilties[flavor]
	if availability == "" {
		return availability, fmt.Errorf("no such flavor: %s", flavor)
	}
	return
}
