// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmCloudantReplicationDataSourceBasic(t *testing.T) {
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	db := fmt.Sprintf("tf_db_%d", acctest.RandIntRange(10, 100))
	partitioned := "false"
	q := "0"
	docID := "doc_id"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCloudantReplicationDataSourceConfigBasic(instanceName, db, partitioned, q, docID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_replication.cloudant_replication", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_replication.cloudant_replication", "cloudant_guid"),
				),
			},
		},
	})
}

func testAccCheckIbmCloudantReplicationDataSourceConfigBasic(instanceName, db string, partitioned string, q string, docID string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "cloudant" {
			is_default=true
	  	}

	  	resource "ibm_resource_instance" "cloudant_instance" {
			name              = "%s"
			service           = "cloudantnosqldb"
			plan              = "standard"
			location          = "us-east"
			resource_group_id = data.ibm_resource_group.cloudant.id
	  	}

		resource "ibm_resource_key" "resource_key" {
			name                 = "pr_key01"
			role                 = "Writer"
			resource_instance_id = ibm_resource_instance.cloudant_instance.id
		}  

		resource "ibm_cloudant_database" "cloudant_database" {
			cloudant_guid = ibm_resource_instance.cloudant_instance.guid
			db = "_replicator"
			partitioned = %s
			q = %s
		}

		resource "ibm_cloudant_replication" "cloudant_replication" {
			doc_id = "%s"
			cloudant_guid = ibm_cloudant_database.cloudant_database.cloudant_guid
			replication_document {
				id = "%s"
				create_target = false
				continuous    = true
				source {
					auth {
				  		iam {
							api_key = concat(ibm_resource_key.resource_key.*.credentials.apikey, [""])[0]
				  		}
					}
					url = concat(ibm_resource_key.resource_key.*.credentials.host, [""])[0]
			  	}
		  
			  	target {
					auth {
				  	iam {
						api_key = concat(ibm_resource_key.resource_key.*.credentials.apikey, [""])[0]
				  	}
				}
				url = concat(ibm_resource_key.resource_key.*.credentials.host, [""])[0]
			  }
			}
		}

		data "ibm_cloudant_replication" "cloudant_replication" {
			doc_id = "%s"
			cloudant_guid = ibm_cloudant_replication.cloudant_replication.cloudant_guid
		}
	`, instanceName, partitioned, q, docID, docID, docID)
}
