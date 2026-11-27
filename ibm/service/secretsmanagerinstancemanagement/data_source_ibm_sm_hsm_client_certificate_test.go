// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

package secretsmanagerinstancemanagement_test

import (
	"testing"
)

func TestAccIbmSmHsmClientCertificateDataSourceBasic(t *testing.T) {
	instanceID := hsmInstanceID(t)
	name := hsmRandName()
	config := hsmDestinationConfig(instanceID) + hsmClientCertBlock(name, "") + `
		data "ibm_sm_hsm_client_certificate" "test" {
			instance_id = ibm_sm_hsm_client_certificate.test.instance_id
			destination_id = ibm_sm_hsm_client_certificate.test.destination_id
			cert_id = ibm_sm_hsm_client_certificate.test.cert_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_client_certificate.test",
		[]string{"id", "instance_id", "destination_id", "cert_id", "name", "provisioning_state", "health_state", "health_as_of", "csr", "csr_cn", "created_at", "created_by"},
		map[string]string{"name": name},
	)
}

func TestAccIbmSmHsmClientCertificateDataSourceAllArgs(t *testing.T) {
	instanceID := hsmInstanceID(t)
	name := hsmRandName()
	description := "tf_description_" + name
	config := hsmDestinationConfig(instanceID) + hsmClientCertBlock(name, description) + `
		data "ibm_sm_hsm_client_certificate" "test" {
			instance_id = ibm_sm_hsm_client_certificate.test.instance_id
			destination_id = ibm_sm_hsm_client_certificate.test.destination_id
			cert_id = ibm_sm_hsm_client_certificate.test.cert_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_client_certificate.test",
		[]string{"id", "instance_id", "destination_id", "cert_id", "name", "description", "provisioning_state", "health_state", "health_as_of", "csr", "csr_cn", "created_at", "created_by"},
		map[string]string{"name": name, "description": description},
	)
}
