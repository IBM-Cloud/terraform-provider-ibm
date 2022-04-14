// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisLogpushJobs_Basic(t *testing.T) {

	name := "MylogpushJob"
	logpull_opt := "timestamps=rfc3339&timestamps=rfc3339"
	data_set := "http_requests"
	freq := "low"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisLogpushJobs_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cis_logpush_job.test", "name", name),
					resource.TestCheckResourceAttr("ibm_cis_logpush_job.test", "enabled", "false"),
					resource.TestCheckResourceAttr("ibm_cis_logpush_job.test", "logpull_options", logpull_opt),
					resource.TestCheckResourceAttr("ibm_cis_logpush_job.test", "dataset", data_set),
					resource.TestCheckResourceAttr("ibm_cis_logpush_job.test", "frequency", freq),
					resource.TestCheckResourceAttrSet("ibm_cis_logpush_job.test", "logdna"),
				),
			},
		},
	})
}
func testAccCheckCisLogpushJobs_basic() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	  resource "ibm_cis_logpush_job" "test" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.domain_id
		name            = "MylogpushJob"
		enabled         = false
		logpull_options = "timestamps=rfc3339&timestamps=rfc3339"
		dataset         = "http_requests"
		frequency       = "low"
		logdna =<<LOG
			{
				"hostname": "travis-kuganes1.sdk.cistest-load.com",
				"ingress_key": "e2f72c7b73a251ca0f7430450b87859e",
				"region": "in-che"
		}
		LOG
	}
`
}
