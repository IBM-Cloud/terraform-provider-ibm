// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisDNSRecordsBatch_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDomainDataSourceConfigBasic1() + testAccCheckIBMCisDNSRecordsBatchConfigPosts(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cis_dns_records_batch.test", "result_posts.#", "2"),
					resource.TestCheckResourceAttr("ibm_cis_dns_records_batch.test", "result_posts.0.type", "A"),
					resource.TestCheckResourceAttr("ibm_cis_dns_records_batch.test", "result_posts.1.type", "TXT"),
				),
			},
		},
	})
}

func TestAccIBMCisDNSRecordsBatch_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDomainDataSourceConfigBasic1() + testAccCheckIBMCisDNSRecordsBatchConfigPosts(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cis_dns_records_batch.test", "result_posts.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMCisDomainDataSourceConfigBasic1() + testAccCheckIBMCisDNSRecordsBatchConfigUpdate(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cis_dns_records_batch.test", "result_posts.#", "2"),
					resource.TestCheckResourceAttr("ibm_cis_dns_records_batch.update", "result_puts.#", "1"),
					resource.TestCheckResourceAttr("ibm_cis_dns_records_batch.update", "result_puts.0.content", "5.6.7.8"),
					resource.TestCheckResourceAttr("ibm_cis_dns_records_batch.update", "result_patches.#", "1"),
					resource.TestCheckResourceAttr("ibm_cis_dns_records_batch.update", "result_deletes.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMCisDNSRecordsBatchConfigPosts() string {
	return `
resource "ibm_cis_dns_records_batch" "test" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.id

  posts {
    name    = "tf-batch-test-a"
    type    = "A"
    content = "1.2.3.4"
    ttl     = 120
    proxied = false
  }

  posts {
    name    = "tf-batch-test-txt"
    type    = "TXT"
    content = "hello from batch"
    ttl     = 300
  }
}
`
}

func testAccCheckIBMCisDNSRecordsBatchConfigUpdate() string {
	return `
resource "ibm_cis_dns_records_batch" "test" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.id

  posts {
    name    = "tf-batch-test-a"
    type    = "A"
    content = "1.2.3.4"
    ttl     = 120
    proxied = false
  }

  posts {
    name    = "tf-batch-test-txt"
    type    = "TXT"
    content = "hello from batch"
    ttl     = 300
  }
}

resource "ibm_cis_dns_records_batch" "update" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.id

  puts {
    id      = ibm_cis_dns_records_batch.test.result_posts[0].id
    name    = "tf-batch-test-a"
    type    = "A"
    content = "5.6.7.8"
    ttl     = 240
    proxied = false
  }

  patches {
    id      = ibm_cis_dns_records_batch.test.result_posts[1].id
    content = "updated via patch"
  }

  deletes {
    id = ibm_cis_dns_records_batch.test.result_posts[0].id
  }
}
`
}
