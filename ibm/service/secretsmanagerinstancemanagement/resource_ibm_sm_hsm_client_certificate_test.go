// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanagerinstancemanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func TestAccIbmSmHsmClientCertificate(t *testing.T) {
	testAccHsmResource(t, hsmResourceSpec{
		resourceType: "ibm_sm_hsm_client_certificate",
		config: func(instanceID, name, description string) string {
			return hsmDestinationConfig(instanceID) + fmt.Sprintf(`
		resource "ibm_sm_hsm_client_certificate" "test" {
			instance_id = ibm_sm_hsm_destination.parent.instance_id
			destination_id = ibm_sm_hsm_destination.parent.destination_id
			name = "%s"
			%s
		}
	`, name, hsmDescriptionLine(description))
		},
		extraChecks: []resource.TestCheckFunc{
			resource.TestCheckResourceAttrSet("ibm_sm_hsm_client_certificate.test", "destination_id"),
			resource.TestCheckResourceAttrSet("ibm_sm_hsm_client_certificate.test", "cert_id"),
			resource.TestCheckResourceAttrSet("ibm_sm_hsm_client_certificate.test", "csr"),
		},
		get: func(client *secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2, parts []string) (*core.DetailedResponse, error) {
			options := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationClientCertificateOptions{}
			options.SetInstanceID(parts[0])
			options.SetDestinationID(parts[1])
			options.SetCertID(parts[2])
			_, response, err := client.GetInstanceHsmDestinationClientCertificate(options)
			return response, err
		},
	}, true)
}
