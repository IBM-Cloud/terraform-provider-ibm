package utils

import (
	"log"
	"net/url"
	"os"
	"reflect"

	"github.com/IBM-Cloud/power-go-client/helpers"
)

// GetNext ...
func GetNext(next interface{}) string {
	if reflect.ValueOf(next).IsNil() {
		return ""
	}

	u, err := url.Parse(reflect.ValueOf(next).Elem().FieldByName("Href").Elem().String())
	if err != nil {
		return ""
	}

	q := u.Query()
	return q.Get("start")
}

// GetEndpoint ...
func GetEndpoint(generation int, regionName string) string {

	switch generation {
	case 1:
		ep := getGCEndpoint(regionName)
		return helpers.EnvFallBack([]string{"IBMCLOUD_IS_API_ENDPOINT"}, ep)
	case 2:
		ep := getNGEndpoint(regionName)
		return helpers.EnvFallBack([]string{"IBMCLOUD_IS_NG_API_ENDPOINT"}, ep)
	}
	ep := getNGEndpoint(regionName)
	return helpers.EnvFallBack([]string{"IBMCLOUD_IS_NG_API_ENDPOINT"}, ep)
}

func getGCEndpoint(regionName string) string {
	if url := os.Getenv("IBMCLOUD_IS_API_ENDPOINT"); url != "" {
		return url
	}
	return regionName + ".iaas.cloud.ibm.com"
}

// For Power-IAAS
func regiontoZone() func(string) string {
	log.Printf("Printing the regiontozone function")
	innerzoneMap := map[string]string{

		"eu-de-1":    "eu-de",
		"eu-de-2":    "eu-de",
		"us-south-1": "us-south",
		"us-south-2": "us-south",
		"us-east":    "us-east",
	}
	return func(key string) string {

		return innerzoneMap[key]
	}

}
func getNGEndpoint(regionName string) string {
	if url := os.Getenv("IBMCLOUD_IS_NG_API_ENDPOINT"); url != "" {
		return url
	}
	return regionName + ".power-iaas.cloud.ibm.com"
}

func GetPowerEndPoint(regionName string) string {
	if url := os.Getenv("IBMCLOUD_IS_NG_API_ENDPOINT"); url != "" {
		return url
	}
	zone := regiontoZone()
	region_name := zone(regionName)
	log.Print("The url to call is %s ", region_name)
	return region_name + ".power-iaas.cloud.ibm.com"

}
