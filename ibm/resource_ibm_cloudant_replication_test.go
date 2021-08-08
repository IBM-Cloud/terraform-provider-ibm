// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

func TestAccIbmCloudantReplicationBasic(t *testing.T) {
	var conf cloudantv1.ReplicationDocument
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	db := fmt.Sprintf("tf_db_%d", acctest.RandIntRange(10, 100))
	partitioned := "false"
	q := "0"
	docID := "doc_id"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmCloudantReplicationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCloudantReplicationConfigBasic(instanceName, db, partitioned, q, docID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCloudantReplicationExists("ibm_cloudant_replication.cloudant_replication", conf),
					resource.TestCheckResourceAttr("ibm_cloudant_replication.cloudant_replication", "doc_id", docID),
				),
			},
		},
	})
}

func TestAccIbmCloudantReplicationAllArgs(t *testing.T) {
	var conf cloudantv1.ReplicationDocument
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	db := fmt.Sprintf("tf_db_%d", acctest.RandIntRange(10, 100))
	partitioned := "false"
	q := "0"
	docID := "doc_id"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmCloudantReplicationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCloudantReplicationConfig(instanceName, db, partitioned, q, docID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCloudantReplicationExists("ibm_cloudant_replication.cloudant_replication", conf),
					resource.TestCheckResourceAttr("ibm_cloudant_replication.cloudant_replication", "doc_id", docID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cloudant_replication.cloudant_replication",
				ImportState:       true,
				ImportStateVerify: true, ImportStateVerifyIgnore: []string{
					"new_edits"},
			},
		},
	})
}

func testAccCheckIbmCloudantReplicationConfigBasic(instanceName, db string, partitioned string, q string, docID string) string {
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
	`, instanceName, partitioned, q, docID, docID)
}

func testAccCheckIbmCloudantReplicationConfig(instanceName, db string, partitioned string, q string, docID string) string {
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
		new_edits = true
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
`, instanceName, partitioned, q, docID, docID)
}

func testAccCheckIbmCloudantReplicationExists(n string, obj cloudantv1.ReplicationDocument) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cloudantClient, err := testAccProvider.Meta().(ClientSession).CloudantV1()
		if err != nil {
			return err
		}

		parts, err := sepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getReplicationDocumentOptions := &cloudantv1.GetReplicationDocumentOptions{}
		getReplicationDocumentOptions.SetDocID(parts[1])

		documentResult, _, err := cloudantClient.GetReplicationDocument(getReplicationDocumentOptions)
		if err != nil {
			return err
		}

		obj = *documentResult
		return nil
	}
}

func testAccCheckIbmCloudantReplicationDestroy(s *terraform.State) error {
	cloudantClient, err := testAccProvider.Meta().(ClientSession).CloudantV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cloudant_replication" {
			continue
		}

		getReplicationDocumentOptions := &cloudantv1.GetReplicationDocumentOptions{}

		parts, err := sepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getReplicationDocumentOptions.SetDocID(parts[1])

		// Try to find the key
		_, _, err = cloudantClient.GetReplicationDocument(getReplicationDocumentOptions)
		if err == nil {
			return fmt.Errorf("cloudant_replication still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
