package contextbasedrestrictions_test

import "os"

const (
	testAccountID = "12ab34cd56ef78ab90cd12ef34ab56cd"
	testZoneID    = "559052eb8f43302824e7ae490c0281eb"
)

func getTestAccountAndZoneID() (string, string) {
	accountID, ok := os.LookupEnv("IBM_IAMACCOUNTID")
	if !ok {
		accountID = testAccountID
	}

	zoneID, ok := os.LookupEnv("IBMCLOUD_CONTEXT_BASED_RESTRICTIONS_ZONE_ID")
	if !ok {
		zoneID = testZoneID
	}
	return accountID, zoneID
}
