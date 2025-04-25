// Copyright IBM Corp. 2024,2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package sdsaas_test

import (
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/sdsaas"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/sds-go-sdk/sdsaasv1"
	"github.com/stretchr/testify/assert"
)

func TestResourceIBMSdsVolumeMappingVolumeReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(sdsaasv1.VolumeReference)
	model.ID = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")

	result, err := sdsaas.ResourceIBMSdsVolumeMappingVolumeReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsVolumeMappingStorageIdentifierToMap(t *testing.T) {
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

	result, err := sdsaas.ResourceIBMSdsVolumeMappingStorageIdentifierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsVolumeMappingGatewayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["ip_address"] = "testString"
		model["port"] = int(22)

		assert.Equal(t, result, model)
	}

	model := new(sdsaasv1.Gateway)
	model.IPAddress = core.StringPtr("testString")
	model.Port = core.Int64Ptr(int64(22))

	result, err := sdsaas.ResourceIBMSdsVolumeMappingGatewayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsVolumeMappingHostReferenceToMap(t *testing.T) {
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

	result, err := sdsaas.ResourceIBMSdsVolumeMappingHostReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsVolumeMappingNamespaceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(1)
		model["uuid"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(sdsaasv1.Namespace)
	model.ID = core.Int64Ptr(int64(1))
	model.UUID = core.StringPtr("testString")

	result, err := sdsaas.ResourceIBMSdsVolumeMappingNamespaceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMSdsVolumeMappingMapToVolumeIdentity(t *testing.T) {
	checkResult := func(result *sdsaasv1.VolumeIdentity) {
		model := new(sdsaasv1.VolumeIdentity)
		model.ID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "testString"

	result, err := sdsaas.ResourceIBMSdsVolumeMappingMapToVolumeIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}
