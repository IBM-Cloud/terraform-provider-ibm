// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package sdsaas_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

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

// NOTE: This test maps a volume to a host. When the test attempts to cleanup the host, the host is still mapped to
// a volume. At this time, the API doesn't support the unmapping between the host and volume at the same time as a delete meaning
// this can't be completed solely through terraform.

// func TestAccIBMSdsHostAllArgs(t *testing.T) {
// 	var conf sdsaasv1.Host
// 	nqn := "nqn.2014-06.org:9345"
// 	name := "terraform-host-name"
// 	nameUpdate := "terraform-host-name-updated"

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { acc.TestAccPreCheck(t) },
// 		Providers:    acc.TestAccProviders,
// 		CheckDestroy: testAccCheckIBMSdsHostDestroy,
// 		Steps: []resource.TestStep{
// 			resource.TestStep{
// 				Config: testAccCheckIBMSdsHostConfig(name, nqn),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					testAccCheckIBMSdsHostExists("ibm_sds_host.sds_host_instance", conf),
// 					resource.TestCheckResourceAttr("ibm_sds_host.sds_host_instance", "name", name),
// 					resource.TestCheckResourceAttr("ibm_sds_host.sds_host_instance", "nqn", nqn),
// 				),
// 			},
// 			resource.TestStep{
// 				Config: testAccCheckIBMSdsHostConfig(nameUpdate, nqn),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr("ibm_sds_host.sds_host_instance", "name", nameUpdate),
// 					resource.TestCheckResourceAttr("ibm_sds_host.sds_host_instance", "nqn", nqn),
// 				),
// 			},
// 			resource.TestStep{
// 				ResourceName:      "ibm_sds_host.sds_host",
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 			},
// 		},
// 	})
// }

func testAccCheckIBMSdsHostConfigBasic(nqn string) string {
	return fmt.Sprintf(`
		resource "ibm_sds_volume" "sds_volume_instance" {
			capacity = 10
			name = "my-volume"
		}
		resource "ibm_sds_host" "sds_host_instance" {
			nqn = "%s"
			name = "my-host"
		}
	`, nqn)
}

func testAccCheckIBMSdsHostConfig(name string, nqn string) string {
	return fmt.Sprintf(`

		output "ibm_sds_volume" {
			value       = [ibm_sds_volume.sds_volume_instance]
			description = "sds_volume resource instance"
		}
		resource "ibm_sds_volume" "sds_volume_instance" {
			capacity = 10
			name = "my-volume"
		}

		resource "ibm_sds_host" "sds_host_instance" {
			name = "%s"
			nqn = "%s"
			volumes {
				volume_name = ibm_sds_volume.sds_volume_instance.name
				volume_id = ibm_sds_volume.sds_volume_instance.id
			}
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

func TestResourceIBMSdsHostVolumeMappingReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		networkInfoReferenceModel := make(map[string]interface{})
		networkInfoReferenceModel["gateway_ip"] = "testString"
		networkInfoReferenceModel["port"] = int(38)

		storageIdentifiersReferenceModel := make(map[string]interface{})
		storageIdentifiersReferenceModel["id"] = "testString"
		storageIdentifiersReferenceModel["namespace_id"] = int(38)
		storageIdentifiersReferenceModel["namespace_uuid"] = "testString"
		storageIdentifiersReferenceModel["network_info"] = []map[string]interface{}{networkInfoReferenceModel}

		model := make(map[string]interface{})
		model["status"] = "testString"
		model["volume_id"] = "testString"
		model["volume_name"] = "testString"
		model["storage_identifiers"] = []map[string]interface{}{storageIdentifiersReferenceModel}

		assert.Equal(t, result, model)
	}

	networkInfoReferenceModel := new(sdsaasv1.NetworkInfoReference)
	networkInfoReferenceModel.GatewayIP = core.StringPtr("testString")
	networkInfoReferenceModel.Port = core.Int64Ptr(int64(38))

	storageIdentifiersReferenceModel := new(sdsaasv1.StorageIdentifiersReference)
	storageIdentifiersReferenceModel.ID = core.StringPtr("testString")
	storageIdentifiersReferenceModel.NamespaceID = core.Int64Ptr(int64(38))
	storageIdentifiersReferenceModel.NamespaceUUID = core.StringPtr("testString")
	storageIdentifiersReferenceModel.NetworkInfo = []sdsaasv1.NetworkInfoReference{*networkInfoReferenceModel}

	model := new(sdsaasv1.VolumeMappingReference)
	model.Status = core.StringPtr("testString")
	model.VolumeID = core.StringPtr("testString")
	model.VolumeName = core.StringPtr("testString")
	model.StorageIdentifiers = storageIdentifiersReferenceModel

	result, err := sdsaas.ResourceIBMSdsHostVolumeMappingReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsHostStorageIdentifiersReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		networkInfoReferenceModel := make(map[string]interface{})
		networkInfoReferenceModel["gateway_ip"] = "testString"
		networkInfoReferenceModel["port"] = int(38)

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["namespace_id"] = int(38)
		model["namespace_uuid"] = "testString"
		model["network_info"] = []map[string]interface{}{networkInfoReferenceModel}

		assert.Equal(t, result, model)
	}

	networkInfoReferenceModel := new(sdsaasv1.NetworkInfoReference)
	networkInfoReferenceModel.GatewayIP = core.StringPtr("testString")
	networkInfoReferenceModel.Port = core.Int64Ptr(int64(38))

	model := new(sdsaasv1.StorageIdentifiersReference)
	model.ID = core.StringPtr("testString")
	model.NamespaceID = core.Int64Ptr(int64(38))
	model.NamespaceUUID = core.StringPtr("testString")
	model.NetworkInfo = []sdsaasv1.NetworkInfoReference{*networkInfoReferenceModel}

	result, err := sdsaas.ResourceIBMSdsHostStorageIdentifiersReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsHostNetworkInfoReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["gateway_ip"] = "testString"
		model["port"] = int(38)

		assert.Equal(t, result, model)
	}

	model := new(sdsaasv1.NetworkInfoReference)
	model.GatewayIP = core.StringPtr("testString")
	model.Port = core.Int64Ptr(int64(38))

	result, err := sdsaas.ResourceIBMSdsHostNetworkInfoReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsHostMapToVolumeMappingIdentity(t *testing.T) {
	checkResult := func(result *sdsaasv1.VolumeMappingIdentity) {
		model := new(sdsaasv1.VolumeMappingIdentity)
		model.VolumeID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["volume_id"] = "testString"

	result, err := sdsaas.ResourceIBMSdsHostMapToVolumeMappingIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}
