// Copyright IBM Corp. 2024,2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package sdsaas_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/sdsaas"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/sds-go-sdk/sdsaasv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMSdsVolumeBasic(t *testing.T) {
	var conf sdsaasv1.Volume
	capacity := fmt.Sprintf("%d", acctest.RandIntRange(1, 5))
	name := "terraform-test-1"
	nameUpdate := "terraform-test-1-updated"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSdsVolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSdsVolumeConfigBasic(capacity, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSdsVolumeExists("ibm_sds_volume.sds_volume_instance", conf),
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "capacity", capacity),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSdsVolumeConfigBasic(capacity, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIBMSdsVolumeAllArgs(t *testing.T) {
	var conf sdsaasv1.Volume
	capacity := fmt.Sprintf("%d", acctest.RandIntRange(1, 5))
	name := "terraform-test-name-1"
	nameUpdate := "terraform-test-name-updated"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSdsVolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSdsVolumeConfig(capacity, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSdsVolumeExists("ibm_sds_volume.sds_volume_instance", conf),
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "capacity", capacity),
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSdsVolumeConfig(capacity, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "capacity", capacity),
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_sds_volume.sds_volume_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSdsVolumeConfigBasic(capacity string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_sds_volume" "sds_volume_instance" {
			capacity = %s
			name = "%s"
		}
	`, capacity, name)
}

func testAccCheckIBMSdsVolumeConfig(capacity string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_sds_volume" "sds_volume_instance" {
			capacity = %s
			name = "%s"
		}
	`, capacity, name)
}

func testAccCheckIBMSdsVolumeExists(n string, obj sdsaasv1.Volume) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		sdsaasClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SdsaasV1()
		if err != nil {
			return err
		}

		volumeOptions := &sdsaasv1.VolumeOptions{}

		volumeOptions.SetVolumeID(rs.Primary.ID)

		// Give time for the volume to fully create before checking its state
		time.Sleep(10 * time.Second)

		volume, _, err := sdsaasClient.Volume(volumeOptions)
		if err != nil {
			return err
		}

		obj = *volume
		return nil
	}
}

func testAccCheckIBMSdsVolumeDestroy(s *terraform.State) error {
	sdsaasClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SdsaasV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sds_volume" {
			continue
		}

		volumeOptions := &sdsaasv1.VolumeOptions{}

		volumeOptions.SetVolumeID(rs.Primary.ID)

		// Give time for the volume to fully delete before checking its state
		time.Sleep(5 * time.Second)

		// Try to find the key
		_, response, err := sdsaasClient.Volume(volumeOptions)

		if err == nil {
			return fmt.Errorf("sds_volume still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for sds_volume (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMSdsVolumeVolumeStatusReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "volume_not_found"
		model["message"] = "Specified resource not found"
		model["more_info"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(sdsaasv1.VolumeStatusReason)
	model.Code = core.StringPtr("volume_not_found")
	model.Message = core.StringPtr("Specified resource not found")
	model.MoreInfo = core.StringPtr("testString")

	result, err := sdsaas.ResourceIBMSdsVolumeVolumeStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsVolumeVolumeMappingToMap(t *testing.T) {
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

	result, err := sdsaas.ResourceIBMSdsVolumeVolumeMappingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsVolumeStorageIdentifierToMap(t *testing.T) {
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

	result, err := sdsaas.ResourceIBMSdsVolumeStorageIdentifierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsVolumeGatewayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["ip_address"] = "testString"
		model["port"] = int(22)

		assert.Equal(t, result, model)
	}

	model := new(sdsaasv1.Gateway)
	model.IPAddress = core.StringPtr("testString")
	model.Port = core.Int64Ptr(int64(22))

	result, err := sdsaas.ResourceIBMSdsVolumeGatewayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsVolumeVolumeReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(sdsaasv1.VolumeReference)
	model.ID = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")

	result, err := sdsaas.ResourceIBMSdsVolumeVolumeReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsVolumeHostReferenceToMap(t *testing.T) {
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

	result, err := sdsaas.ResourceIBMSdsVolumeHostReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsVolumeNamespaceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(1)
		model["uuid"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(sdsaasv1.Namespace)
	model.ID = core.Int64Ptr(int64(1))
	model.UUID = core.StringPtr("testString")

	result, err := sdsaas.ResourceIBMSdsVolumeNamespaceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
