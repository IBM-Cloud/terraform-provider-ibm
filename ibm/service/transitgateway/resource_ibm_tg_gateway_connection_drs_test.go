// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway

import (
	"strings"
	"testing"
)

func TestTransitGatewayConnectionValidatorAllowsDynamicRouteServer(t *testing.T) {
	validator := ResourceIBMTransitGatewayConnectionValidator()
	for _, validationSchema := range validator.Schema {
		if validationSchema.Identifier != tgNetworkType {
			continue
		}

		for _, networkType := range strings.Split(validationSchema.AllowedValues, ",") {
			if strings.TrimSpace(networkType) == "dynamic_route_server" {
				return
			}
		}

		t.Fatalf("dynamic_route_server is missing from network_type allowed values: %q", validationSchema.AllowedValues)
	}

	t.Fatal("network_type validator is missing")
}

func TestTransitGatewayConnectionDynamicRouteServerSchema(t *testing.T) {
	connectionSchema := ResourceIBMTransitGatewayConnection().Schema

	if !connectionSchema[tgNetworkId].Optional || !connectionSchema[tgNetworkId].ForceNew {
		t.Fatal("network_id must remain an optional, ForceNew input for dynamic_route_server connections")
	}
	if !connectionSchema[tgCidr].Optional || !connectionSchema[tgCidr].ForceNew {
		t.Fatal("cidr must be an optional, ForceNew input for dynamic_route_server connections")
	}
}
