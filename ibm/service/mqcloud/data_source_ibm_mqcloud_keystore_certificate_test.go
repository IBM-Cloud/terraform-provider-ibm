// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.95.2-120e65bc-20240924-152329
 */

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/mqcloud"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmMqcloudKeystoreCertificateDataSourceBasic(t *testing.T) {
	keyStoreCertificateDetailsServiceInstanceGuid := acc.MqcloudDeploymentID
	keyStoreCertificateDetailsQueueManagerID := acc.MqcloudQueueManagerID
	keyStoreCertificateDetailsLabel := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	keyStoreCertificateDetailsCertificateFile := acc.MqcloudKSCertFilePath

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acc.TestAccPreCheckMqcloud(t)
		},
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudKeystoreCertificateDataSourceConfigBasic(keyStoreCertificateDetailsServiceInstanceGuid, keyStoreCertificateDetailsQueueManagerID, keyStoreCertificateDetailsLabel, keyStoreCertificateDetailsCertificateFile),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance", "queue_manager_id"),
				),
			},
		},
	})
}

func testAccCheckIbmMqcloudKeystoreCertificateDataSourceConfigBasic(keyStoreCertificateDetailsServiceInstanceGuid string, keyStoreCertificateDetailsQueueManagerID string, keyStoreCertificateDetailsLabel string, keyStoreCertificateDetailsCertificateFile string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
			service_instance_guid = "%s"
			queue_manager_id = "%s"
			label = "%s"
			certificate_file = filebase64("%s")
		}

		data "ibm_mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
			service_instance_guid = ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance.service_instance_guid
			queue_manager_id = ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance.queue_manager_id
			label = ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance.label
		}
	`, keyStoreCertificateDetailsServiceInstanceGuid, keyStoreCertificateDetailsQueueManagerID, keyStoreCertificateDetailsLabel, keyStoreCertificateDetailsCertificateFile)
}

func TestDataSourceIbmMqcloudKeystoreCertificateKeyStoreCertificateDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		channelDetailsModel := make(map[string]interface{})
		channelDetailsModel["name"] = "CLOUD.APP.SVRCONN"

		channelsDetailsModel := make(map[string]interface{})
		channelsDetailsModel["channels"] = []map[string]interface{}{channelDetailsModel}

		certificateConfigurationModel := make(map[string]interface{})
		certificateConfigurationModel["ams"] = []map[string]interface{}{channelsDetailsModel}

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["label"] = "testString"
		model["certificate_type"] = "key_store"
		model["fingerprint_sha256"] = "testString"
		model["subject_dn"] = "testString"
		model["subject_cn"] = "testString"
		model["issuer_dn"] = "testString"
		model["issuer_cn"] = "testString"
		model["issued"] = "2019-01-01T12:00:00.000Z"
		model["expiry"] = "2019-01-01T12:00:00.000Z"
		model["is_default"] = true
		model["dns_names_total_count"] = int(38)
		model["dns_names"] = []string{"testString"}
		model["href"] = "testString"
		model["config"] = []map[string]interface{}{certificateConfigurationModel}

		assert.Equal(t, result, model)
	}

	channelDetailsModel := new(mqcloudv1.ChannelDetails)
	channelDetailsModel.Name = core.StringPtr("CLOUD.APP.SVRCONN")

	channelsDetailsModel := new(mqcloudv1.ChannelsDetails)
	channelsDetailsModel.Channels = []mqcloudv1.ChannelDetails{*channelDetailsModel}

	certificateConfigurationModel := new(mqcloudv1.CertificateConfiguration)
	certificateConfigurationModel.Ams = channelsDetailsModel

	model := new(mqcloudv1.KeyStoreCertificateDetails)
	model.ID = core.StringPtr("testString")
	model.Label = core.StringPtr("testString")
	model.CertificateType = core.StringPtr("key_store")
	model.FingerprintSha256 = core.StringPtr("testString")
	model.SubjectDn = core.StringPtr("testString")
	model.SubjectCn = core.StringPtr("testString")
	model.IssuerDn = core.StringPtr("testString")
	model.IssuerCn = core.StringPtr("testString")
	model.Issued = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.Expiry = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.IsDefault = core.BoolPtr(true)
	model.DnsNamesTotalCount = core.Int64Ptr(int64(38))
	model.DnsNames = []string{"testString"}
	model.Href = core.StringPtr("testString")
	model.Config = certificateConfigurationModel

	result, err := mqcloud.DataSourceIbmMqcloudKeystoreCertificateKeyStoreCertificateDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmMqcloudKeystoreCertificateCertificateConfigurationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		channelDetailsModel := make(map[string]interface{})
		channelDetailsModel["name"] = "CLOUD.APP.SVRCONN"

		channelsDetailsModel := make(map[string]interface{})
		channelsDetailsModel["channels"] = []map[string]interface{}{channelDetailsModel}

		model := make(map[string]interface{})
		model["ams"] = []map[string]interface{}{channelsDetailsModel}

		assert.Equal(t, result, model)
	}

	channelDetailsModel := new(mqcloudv1.ChannelDetails)
	channelDetailsModel.Name = core.StringPtr("CLOUD.APP.SVRCONN")

	channelsDetailsModel := new(mqcloudv1.ChannelsDetails)
	channelsDetailsModel.Channels = []mqcloudv1.ChannelDetails{*channelDetailsModel}

	model := new(mqcloudv1.CertificateConfiguration)
	model.Ams = channelsDetailsModel

	result, err := mqcloud.DataSourceIbmMqcloudKeystoreCertificateCertificateConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmMqcloudKeystoreCertificateChannelsDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		channelDetailsModel := make(map[string]interface{})
		channelDetailsModel["name"] = "CLOUD.APP.SVRCONN"

		model := make(map[string]interface{})
		model["channels"] = []map[string]interface{}{channelDetailsModel}

		assert.Equal(t, result, model)
	}

	channelDetailsModel := new(mqcloudv1.ChannelDetails)
	channelDetailsModel.Name = core.StringPtr("CLOUD.APP.SVRCONN")

	model := new(mqcloudv1.ChannelsDetails)
	model.Channels = []mqcloudv1.ChannelDetails{*channelDetailsModel}

	result, err := mqcloud.DataSourceIbmMqcloudKeystoreCertificateChannelsDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmMqcloudKeystoreCertificateChannelDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(mqcloudv1.ChannelDetails)
	model.Name = core.StringPtr("testString")

	result, err := mqcloud.DataSourceIbmMqcloudKeystoreCertificateChannelDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
