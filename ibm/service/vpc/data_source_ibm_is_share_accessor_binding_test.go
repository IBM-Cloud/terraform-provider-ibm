// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.90.0-5aad763d-20240506-203857
 */

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsShareAccessorBindingDataSourceBasic(t *testing.T) {
	subnetName := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tf-vpc-name-%d", acctest.RandIntRange(10, 100))
	shareName := fmt.Sprintf("tf-share-%d", acctest.RandIntRange(10, 100))
	shareName1 := fmt.Sprintf("tf-share1-%d", acctest.RandIntRange(10, 100))
	tEMode1 := "user_managed"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsShareAccessorBindingDataSourceConfigBasic(vpcname, subnetName, tEMode1, shareName, shareName1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_share_accessor_binding.is_share_accessor_binding_instance", "share"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_accessor_binding.is_share_accessor_binding_instance", "accessor.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_accessor_binding.is_share_accessor_binding_instance", "accessor.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_accessor_binding.is_share_accessor_binding_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_accessor_binding.is_share_accessor_binding_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_accessor_binding.is_share_accessor_binding_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_accessor_binding.is_share_accessor_binding_instance", "resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsShareAccessorBindingDataSourceConfigBasic(vpcName, sname, tEMode, shareName, shareName1 string) string {
	return testAccCheckIbmIsShareConfigOriginShareConfig(vpcName, sname, tEMode, shareName, shareName1) + fmt.Sprintf(`
		data "ibm_is_share_accessor_bindings" "bindings" {
			depends_on = [ibm_is_share.is_share_accessor]
			share = ibm_is_share.is_share.id
		}
		data "ibm_is_share_accessor_binding" "is_share_accessor_binding_instance" {
			share = ibm_is_share.is_share.id
			accessor_binding = data.ibm_is_share_accessor_bindings.bindings.accessor_bindings.0.id
		}
	`)
}

func TestDataSourceIBMIsShareAccessorBindingShareAccessorBindingAccessorToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		shareReferenceDeletedModel := make(map[string]interface{})
		shareReferenceDeletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		accountReferenceModel := make(map[string]interface{})
		accountReferenceModel["id"] = "bb1b52262f7441a586f49068482f1e60"
		accountReferenceModel["resource_type"] = "account"

		regionReferenceModel := make(map[string]interface{})
		regionReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south"
		regionReferenceModel["name"] = "us-south"

		shareRemoteModel := make(map[string]interface{})
		shareRemoteModel["account"] = []map[string]interface{}{accountReferenceModel}
		shareRemoteModel["region"] = []map[string]interface{}{regionReferenceModel}

		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::share:0fe9e5d8-0a4d-4818-96ec-e99708644a58"
		model["deleted"] = []map[string]interface{}{shareReferenceDeletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/shares/0fe9e5d8-0a4d-4818-96ec-e99708644a58"
		model["id"] = "0fe9e5d8-0a4d-4818-96ec-e99708644a58"
		model["name"] = "my-share"
		model["remote"] = []map[string]interface{}{shareRemoteModel}
		model["resource_type"] = "share"

		assert.Equal(t, result, model)
	}

	shareReferenceDeletedModel := new(vpcv1.Deleted)
	shareReferenceDeletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	accountReferenceModel := new(vpcv1.AccountReference)
	accountReferenceModel.ID = core.StringPtr("bb1b52262f7441a586f49068482f1e60")
	accountReferenceModel.ResourceType = core.StringPtr("account")

	regionReferenceModel := new(vpcv1.RegionReference)
	regionReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south")
	regionReferenceModel.Name = core.StringPtr("us-south")

	shareRemoteModel := new(vpcv1.ShareRemote)
	shareRemoteModel.Account = accountReferenceModel
	shareRemoteModel.Region = regionReferenceModel

	model := new(vpcv1.ShareAccessorBindingAccessor)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::share:0fe9e5d8-0a4d-4818-96ec-e99708644a58")
	model.Deleted = shareReferenceDeletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/shares/0fe9e5d8-0a4d-4818-96ec-e99708644a58")
	model.ID = core.StringPtr("0fe9e5d8-0a4d-4818-96ec-e99708644a58")
	model.Name = core.StringPtr("my-share")
	model.Remote = shareRemoteModel
	model.ResourceType = core.StringPtr("share")

	result, err := vpc.DataSourceIBMIsShareAccessorBindingShareAccessorBindingAccessorToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareAccessorBindingShareReferenceDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsShareAccessorBindingShareReferenceDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareAccessorBindingShareRemoteToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		accountReferenceModel := make(map[string]interface{})
		accountReferenceModel["id"] = "bb1b52262f7441a586f49068482f1e60"
		accountReferenceModel["resource_type"] = "account"

		regionReferenceModel := make(map[string]interface{})
		regionReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south"
		regionReferenceModel["name"] = "us-south"

		model := make(map[string]interface{})
		model["account"] = []map[string]interface{}{accountReferenceModel}
		model["region"] = []map[string]interface{}{regionReferenceModel}

		assert.Equal(t, result, model)
	}

	accountReferenceModel := new(vpcv1.AccountReference)
	accountReferenceModel.ID = core.StringPtr("bb1b52262f7441a586f49068482f1e60")
	accountReferenceModel.ResourceType = core.StringPtr("account")

	regionReferenceModel := new(vpcv1.RegionReference)
	regionReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south")
	regionReferenceModel.Name = core.StringPtr("us-south")

	model := new(vpcv1.ShareRemote)
	model.Account = accountReferenceModel
	model.Region = regionReferenceModel

	result, err := vpc.DataSourceIBMIsShareAccessorBindingShareRemoteToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareAccessorBindingAccountReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "bb1b52262f7441a586f49068482f1e60"
		model["resource_type"] = "account"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.AccountReference)
	model.ID = core.StringPtr("bb1b52262f7441a586f49068482f1e60")
	model.ResourceType = core.StringPtr("account")

	result, err := vpc.DataSourceIBMIsShareAccessorBindingAccountReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareAccessorBindingRegionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south"
		model["name"] = "us-south"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.RegionReference)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south")
	model.Name = core.StringPtr("us-south")

	result, err := vpc.DataSourceIBMIsShareAccessorBindingRegionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareAccessorBindingShareAccessorBindingAccessorShareReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		shareReferenceDeletedModel := make(map[string]interface{})
		shareReferenceDeletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		accountReferenceModel := make(map[string]interface{})
		accountReferenceModel["id"] = "bb1b52262f7441a586f49068482f1e60"
		accountReferenceModel["resource_type"] = "account"

		regionReferenceModel := make(map[string]interface{})
		regionReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south"
		regionReferenceModel["name"] = "us-south"

		shareRemoteModel := make(map[string]interface{})
		shareRemoteModel["account"] = []map[string]interface{}{accountReferenceModel}
		shareRemoteModel["region"] = []map[string]interface{}{regionReferenceModel}

		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::share:0fe9e5d8-0a4d-4818-96ec-e99708644a58"
		model["deleted"] = []map[string]interface{}{shareReferenceDeletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/shares/0fe9e5d8-0a4d-4818-96ec-e99708644a58"
		model["id"] = "0fe9e5d8-0a4d-4818-96ec-e99708644a58"
		model["name"] = "my-share"
		model["remote"] = []map[string]interface{}{shareRemoteModel}
		model["resource_type"] = "share"

		assert.Equal(t, result, model)
	}

	shareReferenceDeletedModel := new(vpcv1.Deleted)
	shareReferenceDeletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	accountReferenceModel := new(vpcv1.AccountReference)
	accountReferenceModel.ID = core.StringPtr("bb1b52262f7441a586f49068482f1e60")
	accountReferenceModel.ResourceType = core.StringPtr("account")

	regionReferenceModel := new(vpcv1.RegionReference)
	regionReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south")
	regionReferenceModel.Name = core.StringPtr("us-south")

	shareRemoteModel := new(vpcv1.ShareRemote)
	shareRemoteModel.Account = accountReferenceModel
	shareRemoteModel.Region = regionReferenceModel

	model := new(vpcv1.ShareAccessorBindingAccessorShareReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::share:0fe9e5d8-0a4d-4818-96ec-e99708644a58")
	model.Deleted = shareReferenceDeletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/shares/0fe9e5d8-0a4d-4818-96ec-e99708644a58")
	model.ID = core.StringPtr("0fe9e5d8-0a4d-4818-96ec-e99708644a58")
	model.Name = core.StringPtr("my-share")
	model.Remote = shareRemoteModel
	model.ResourceType = core.StringPtr("share")

	result, err := vpc.DataSourceIBMIsShareAccessorBindingShareAccessorBindingAccessorShareReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsShareAccessorBindingShareAccessorBindingAccessorWatsonxMachineLearningReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:pm-20:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34:6500f05d-a5b5-4ecf-91ba-0d12b9dee607"
		model["resource_type"] = "watsonx_machine_learning"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ShareAccessorBindingAccessorWatsonxMachineLearningReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:pm-20:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34:6500f05d-a5b5-4ecf-91ba-0d12b9dee607")
	model.ResourceType = core.StringPtr("watsonx_machine_learning")

	result, err := vpc.DataSourceIBMIsShareAccessorBindingShareAccessorBindingAccessorWatsonxMachineLearningReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
