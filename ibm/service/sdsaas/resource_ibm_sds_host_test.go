// Copyright IBM Corp. 2024,2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package sdsaas_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/sdsaas"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/sds-go-sdk/sdsaasv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMSdsHostBasic(t *testing.T) {
	var conf sdsaasv1.Host
	nqn := "nqn.2014-06.org:9345"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSdsHostDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSdsHostConfigBasic(nqn),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSdsHostExists("ibm_sds_host.sds_host_instance", conf),
					resource.TestCheckResourceAttr("ibm_sds_host.sds_host_instance", "nqn", nqn),
				),
			},
		},
	})
}

func TestAccIBMSdsHostAllArgs(t *testing.T) {
	var conf sdsaasv1.Host

	nqn := "nqn.2014-06.org:9345"
	name := "terraform-host-name"
	nameUpdate := "terraform-host-name-updated"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSdsHostDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSdsHostConfig(name, nqn),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSdsHostExists("ibm_sds_host.sds_host_instance", conf),
					resource.TestCheckResourceAttr("ibm_sds_host.sds_host_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_sds_host.sds_host_instance", "nqn", nqn),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSdsHostConfig(nameUpdate, nqn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_sds_host.sds_host_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_sds_host.sds_host_instance", "nqn", nqn),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_sds_host.sds_host_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSdsHostConfigBasic(nqn string) string {
	return fmt.Sprintf(`
		resource "ibm_sds_host" "sds_host_instance" {
			nqn = "%s"
			name = "my-host"
		}
	`, nqn)
}

func testAccCheckIBMSdsHostConfig(name string, nqn string) string {
	return fmt.Sprintf(`

		resource "ibm_sds_host" "sds_host_instance" {
			name = "%s"
			nqn = "%s"
		}

	`, name, nqn)
}

func testAccCheckIBMSdsHostExists(n string, obj sdsaasv1.Host) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		sdsaasClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SdsaasV1()
		if err != nil {
			return err
		}

		hostOptions := &sdsaasv1.HostOptions{}

		hostOptions.SetHostID(rs.Primary.ID)

		host, _, err := sdsaasClient.Host(hostOptions)
		if err != nil {
			return err
		}

		obj = *host
		return nil
	}
}

func testAccCheckIBMSdsHostDestroy(s *terraform.State) error {
	sdsaasClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SdsaasV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sds_host" {
			continue
		}

		hostOptions := &sdsaasv1.HostOptions{}

		hostOptions.SetHostID(rs.Primary.ID)

		// Try to find the key
		_, response, err := sdsaasClient.Host(hostOptions)

		if err == nil {
			return fmt.Errorf("sds_host still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for sds_host (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMSdsHostVolumeMappingToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		gatewayModel := make(map[string]interface{})
		gatewayModel["ip_address"] = "testString"
		gatewayModel["port"] = int(22)

		storageIdentifierModel := make(map[string]interface{})
		storageIdentifierModel["subsystem_nqn"] = "nqn.2014-06.org:1234"
		storageIdentifierModel["namespace_id"] = int(1)
		storageIdentifierModel["namespace_uuid"] = "testString"
		storageIdentifierModel["gateways"] = []map[string]interface{}{gatewayModel}

		volumeReferenceModel := make(map[string]interface{})
		volumeReferenceModel["id"] = "testString"
		volumeReferenceModel["name"] = "testString"

		hostReferenceModel := make(map[string]interface{})
		hostReferenceModel["id"] = "testString"
		hostReferenceModel["name"] = "testString"
		hostReferenceModel["nqn"] = "testString"

		namespaceModel := make(map[string]interface{})
		namespaceModel["id"] = int(1)
		namespaceModel["uuid"] = "testString"

		model := make(map[string]interface{})
		model["status"] = "mapped"
		model["storage_identifier"] = []map[string]interface{}{storageIdentifierModel}
		model["href"] = "testString"
		model["id"] = "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45"
		model["volume"] = []map[string]interface{}{volumeReferenceModel}
		model["host"] = []map[string]interface{}{hostReferenceModel}
		model["subsystem_nqn"] = "nqn.2014-06.org:1234"
		model["namespace"] = []map[string]interface{}{namespaceModel}
		model["gateways"] = []map[string]interface{}{gatewayModel}

		assert.Equal(t, result, model)
	}

	gatewayModel := new(sdsaasv1.Gateway)
	gatewayModel.IPAddress = core.StringPtr("testString")
	gatewayModel.Port = core.Int64Ptr(int64(22))

	storageIdentifierModel := new(sdsaasv1.StorageIdentifier)
	storageIdentifierModel.SubsystemNqn = core.StringPtr("nqn.2014-06.org:1234")
	storageIdentifierModel.NamespaceID = core.Int64Ptr(int64(1))
	storageIdentifierModel.NamespaceUUID = core.StringPtr("testString")
	storageIdentifierModel.Gateways = []sdsaasv1.Gateway{*gatewayModel}

	volumeReferenceModel := new(sdsaasv1.VolumeReference)
	volumeReferenceModel.ID = core.StringPtr("testString")
	volumeReferenceModel.Name = core.StringPtr("testString")

	hostReferenceModel := new(sdsaasv1.HostReference)
	hostReferenceModel.ID = core.StringPtr("testString")
	hostReferenceModel.Name = core.StringPtr("testString")
	hostReferenceModel.Nqn = core.StringPtr("testString")

	namespaceModel := new(sdsaasv1.Namespace)
	namespaceModel.ID = core.Int64Ptr(int64(1))
	namespaceModel.UUID = core.StringPtr("testString")

	model := new(sdsaasv1.VolumeMapping)
	model.Status = core.StringPtr("mapped")
	model.StorageIdentifier = storageIdentifierModel
	model.Href = core.StringPtr("testString")
	model.ID = core.StringPtr("1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45")
	model.Volume = volumeReferenceModel
	model.Host = hostReferenceModel
	model.SubsystemNqn = core.StringPtr("nqn.2014-06.org:1234")
	model.Namespace = namespaceModel
	model.Gateways = []sdsaasv1.Gateway{*gatewayModel}

	result, err := sdsaas.ResourceIBMSdsHostVolumeMappingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsHostStorageIdentifierToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		gatewayModel := make(map[string]interface{})
		gatewayModel["ip_address"] = "testString"
		gatewayModel["port"] = int(22)

		model := make(map[string]interface{})
		model["subsystem_nqn"] = "nqn.2014-06.org:1234"
		model["namespace_id"] = int(1)
		model["namespace_uuid"] = "testString"
		model["gateways"] = []map[string]interface{}{gatewayModel}

		assert.Equal(t, result, model)
	}

	gatewayModel := new(sdsaasv1.Gateway)
	gatewayModel.IPAddress = core.StringPtr("testString")
	gatewayModel.Port = core.Int64Ptr(int64(22))

	model := new(sdsaasv1.StorageIdentifier)
	model.SubsystemNqn = core.StringPtr("nqn.2014-06.org:1234")
	model.NamespaceID = core.Int64Ptr(int64(1))
	model.NamespaceUUID = core.StringPtr("testString")
	model.Gateways = []sdsaasv1.Gateway{*gatewayModel}

	result, err := sdsaas.ResourceIBMSdsHostStorageIdentifierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsHostGatewayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["ip_address"] = "testString"
		model["port"] = int(22)

		assert.Equal(t, result, model)
	}

	model := new(sdsaasv1.Gateway)
	model.IPAddress = core.StringPtr("testString")
	model.Port = core.Int64Ptr(int64(22))

	result, err := sdsaas.ResourceIBMSdsHostGatewayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsHostVolumeReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(sdsaasv1.VolumeReference)
	model.ID = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")

	result, err := sdsaas.ResourceIBMSdsHostVolumeReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsHostHostReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["name"] = "testString"
		model["nqn"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(sdsaasv1.HostReference)
	model.ID = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.Nqn = core.StringPtr("testString")

	result, err := sdsaas.ResourceIBMSdsHostHostReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsHostNamespaceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(1)
		model["uuid"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(sdsaasv1.Namespace)
	model.ID = core.Int64Ptr(int64(1))
	model.UUID = core.StringPtr("testString")

	result, err := sdsaas.ResourceIBMSdsHostNamespaceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsHostMapToVolumeMappingPrototype(t *testing.T) {
	checkResult := func(result *sdsaasv1.VolumeMappingPrototype) {
		volumeIdentityModel := new(sdsaasv1.VolumeIdentity)
		volumeIdentityModel.ID = core.StringPtr("testString")

		model := new(sdsaasv1.VolumeMappingPrototype)
		model.Volume = volumeIdentityModel

		assert.Equal(t, result, model)
	}

	volumeIdentityModel := make(map[string]interface{})
	volumeIdentityModel["id"] = "testString"

	model := make(map[string]interface{})
	model["volume"] = []interface{}{volumeIdentityModel}

	result, err := sdsaas.ResourceIBMSdsHostMapToVolumeMappingPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsHostMapToVolumeIdentity(t *testing.T) {
	checkResult := func(result *sdsaasv1.VolumeIdentity) {
		model := new(sdsaasv1.VolumeIdentity)
		model.ID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "testString"

	result, err := sdsaas.ResourceIBMSdsHostMapToVolumeIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}
