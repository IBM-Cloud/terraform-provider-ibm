// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

package secretsmanagerinstancemanagement_test

import (
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/secretsmanagerinstancemanagement"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmSmHsmClientCertificatesDataSourceBasic(t *testing.T) {
	instanceID := hsmInstanceID(t)
	name := hsmRandName()
	config := hsmDestinationConfig(instanceID) + hsmClientCertBlock(name, "") + `
		data "ibm_sm_hsm_client_certificates" "test" {
			instance_id = ibm_sm_hsm_client_certificate.test.instance_id
			destination_id = ibm_sm_hsm_client_certificate.test.destination_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_client_certificates.test",
		[]string{"id", "instance_id", "destination_id", "client_certificates.#"},
		map[string]string{"client_certificates.0.name": name},
	)
}

func TestAccIbmSmHsmClientCertificatesDataSourceAllArgs(t *testing.T) {
	instanceID := hsmInstanceID(t)
	name := hsmRandName()
	description := "tf_description_" + name
	config := hsmDestinationConfig(instanceID) + hsmClientCertBlock(name, description) + `
		data "ibm_sm_hsm_client_certificates" "test" {
			instance_id = ibm_sm_hsm_client_certificate.test.instance_id
			destination_id = ibm_sm_hsm_client_certificate.test.destination_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_client_certificates.test",
		[]string{"id", "instance_id", "destination_id", "client_certificates.#", "client_certificates.0.id", "client_certificates.0.href", "client_certificates.0.provisioning_state", "client_certificates.0.health_state", "client_certificates.0.health_as_of", "client_certificates.0.csr", "client_certificates.0.csr_cn", "client_certificates.0.created_at", "client_certificates.0.created_by"},
		map[string]string{"client_certificates.0.name": name, "client_certificates.0.description": description},
	)
}

func TestDataSourceIbmSmHsmClientCertificatesClientCertificateToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "7c9e6679-7425-40de-944b-e07fc1f90ae7"
		model["href"] = "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7"
		model["name"] = "hsm-client-cert-prod"
		model["description"] = "Primary client cert for prod"
		model["status"] = "testString"
		model["provisioning_state"] = "active"
		model["health_state"] = "valid"
		model["health_as_of"] = "2026-07-09T15:30:00.000Z"
		model["csr"] = "-----BEGIN CERTIFICATE REQUEST-----\nMIICvDCCAaQCAQAwdzEL...\n-----END CERTIFICATE REQUEST-----"
		model["csr_cn"] = "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
		model["signed_cert"] = "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----"
		model["expires_at"] = "2028-07-09T10:00:00.000Z"
		model["created_at"] = "2026-07-09T10:00:00.000Z"
		model["updated_at"] = "2026-07-09T11:00:00.000Z"
		model["created_by"] = "iam-ServiceId-abc123"
		model["updated_by"] = "iam-ServiceId-abc123"

		assert.Equal(t, result, model)
	}

	model := new(secretsmanagerinstancemanagementv2.ClientCertificate)
	model.ID = core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")
	model.Href = core.StringPtr("https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7")
	model.Name = core.StringPtr("hsm-client-cert-prod")
	model.Description = core.StringPtr("Primary client cert for prod")
	model.Status = core.StringPtr("testString")
	model.ProvisioningState = core.StringPtr("active")
	model.HealthState = core.StringPtr("valid")
	model.HealthAsOf = CreateMockDateTime("2026-07-09T15:30:00.000Z")
	model.Csr = core.StringPtr("-----BEGIN CERTIFICATE REQUEST-----\nMIICvDCCAaQCAQAwdzEL...\n-----END CERTIFICATE REQUEST-----")
	model.CsrCn = core.StringPtr("a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	model.SignedCert = core.StringPtr("-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----")
	model.ExpiresAt = CreateMockDateTime("2028-07-09T10:00:00.000Z")
	model.CreatedAt = CreateMockDateTime("2026-07-09T10:00:00.000Z")
	model.UpdatedAt = CreateMockDateTime("2026-07-09T11:00:00.000Z")
	model.CreatedBy = core.StringPtr("iam-ServiceId-abc123")
	model.UpdatedBy = core.StringPtr("iam-ServiceId-abc123")

	result, err := secretsmanagerinstancemanagement.DataSourceIbmSmHsmClientCertificatesClientCertificateToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
