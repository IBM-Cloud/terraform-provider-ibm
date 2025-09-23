package atracker

import (
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
)

const (
	REDACTED_TEXT = "REDACTED"
)

func getAtrackerClients(meta interface{}) (
	atrackerClientv2 *atrackerv2.AtrackerV2, err error) {
	atrackerClientv2, err = meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		return
	}

	_, err = meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return
	}

	return atrackerClientv2, nil
}
