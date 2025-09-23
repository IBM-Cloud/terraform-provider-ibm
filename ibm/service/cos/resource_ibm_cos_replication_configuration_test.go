package cos_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Replication features
func TestAccIBMCosBucket_Bucket_Replication(t *testing.T) {

	accountID := acc.IBM_AccountID_REPL
	cosServiceNameSrc := fmt.Sprintf("cos_instance_src_%d", acctest.RandIntRange(10, 100))
	cosServiceNameDest := fmt.Sprintf("cos_instance_dest_%d", acctest.RandIntRange(10, 100))
	bucketNameSrc := fmt.Sprintf("terraform-testacc-src-%d", acctest.RandIntRange(10, 100))
	bucketNameDest := fmt.Sprintf("terraform-testacc-dest-%d", acctest.RandIntRange(10, 100))
	bucketRegionSrc := "us-south"
	bucketRegionDest := "us-south"
	bucketClassSrc := "standard"
	bucketClassDest := "standard"
	bucketRegionType := "region_location"
	ruleId := "my-rule-id-bucket-replication"
	enable := true
	prefix := ""
	priority := 1
	deletemarker_replication_status := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_replication(accountID, cosServiceNameSrc, cosServiceNameDest, bucketNameSrc, bucketNameDest, bucketRegionType, bucketRegionSrc, bucketRegionDest, bucketClassSrc, bucketClassDest, ruleId, enable, priority, deletemarker_replication_status, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_bucket.cos_bucket_source", "bucket_name", bucketNameSrc),
					resource.TestCheckResourceAttr("ibm_cos_bucket.cos_bucket_destination", "bucket_name", bucketNameDest),
					resource.TestCheckResourceAttr("ibm_cos_bucket.cos_bucket_source", "storage_class", bucketClassSrc),
					resource.TestCheckResourceAttr("ibm_cos_bucket.cos_bucket_destination", "storage_class", bucketClassDest),
					resource.TestCheckResourceAttr("ibm_cos_bucket.cos_bucket_source", "region_location", bucketRegionSrc),
					resource.TestCheckResourceAttr("ibm_cos_bucket.cos_bucket_destination", "region_location", bucketRegionDest),
					resource.TestCheckResourceAttr("ibm_cos_bucket.cos_bucket_source", "object_versioning.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.cos_bucket_destination", "object_versioning.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_replication_rule.cos_bucket_repl", "replication_rule.#", "1"),
				),
			},
		},
	})
}

// create cos instance  & buckets for source and destination and enable replication rule after iam authorization policy set on resource attributes..
func testAccCheckIBMCosBucket_replication(accountID, cosServiceNameSrc string, cosServiceNameDest string, bucketNameSrc string, bucketNameDest string, regiontype string, regionSrc string, regionDest string, storageClassSrc string, storageClassDest string, ruleId string, enable bool, priority int, deletemarker_replication_status bool, prefix string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
	}
	resource "ibm_resource_instance" "cos_instance_source" {
		name              = "%s"
		resource_group_id = data.ibm_resource_group.cos_group.id
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
	}

	resource "ibm_cos_bucket" "cos_bucket_source" {
		bucket_name           = "%s"
		resource_instance_id  = ibm_resource_instance.cos_instance_source.id
		region_location   = "%s"
		storage_class         = "%s"
		object_versioning {
		  enable  = true
		}
	}
	resource "ibm_resource_instance" "cos_instance_destination" {
		name              = "%s"
		resource_group_id = data.ibm_resource_group.cos_group.id
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
	}

	resource "ibm_cos_bucket" "cos_bucket_destination" {
		bucket_name           = "%s"
		resource_instance_id  = ibm_resource_instance.cos_instance_destination.id
		region_location   = "%s"
		storage_class         = "%s"
		object_versioning {
		  enable  = true
		}
	}

	resource "ibm_iam_authorization_policy" "policy" {
		roles                  = [
			"Writer",
		]
		subject_attributes {
		  name  = "accountId"
		  value = "%s"
		}
		subject_attributes {
		  name  = "serviceName"
		  value = "cloud-object-storage"
		}
		subject_attributes {
		  name  = "serviceInstance"
		  value = ibm_resource_instance.cos_instance_source.guid
		}
		subject_attributes {
		  name  = "resource"
		  value = ibm_cos_bucket.cos_bucket_source.bucket_name
		}
		subject_attributes {
		  name  = "resourceType"
		  value = "bucket"
		}
		resource_attributes {
		  name     = "accountId"
		  operator = "stringEquals"
		  value    = "%s"
		}
		resource_attributes {
		  name     = "serviceName"
		  operator = "stringEquals"
		  value    = "cloud-object-storage"
		}
		resource_attributes { 
		  name  =  "serviceInstance"
		  operator = "stringEquals"
		  value =  ibm_resource_instance.cos_instance_destination.guid
		}
		resource_attributes { 
		  name  =  "resource"
		  operator = "stringEquals"
		  value =  ibm_cos_bucket.cos_bucket_destination.bucket_name
		}
		resource_attributes { 
		  name  =  "resourceType"
		  operator = "stringEquals"
		  value =  "bucket" 
		}
	}
	resource "ibm_cos_bucket_replication_rule" "cos_bucket_repl" {
		depends_on = [
			ibm_iam_authorization_policy.policy
		]
		bucket_crn	    = ibm_cos_bucket.cos_bucket_source.crn
		bucket_location = ibm_cos_bucket.cos_bucket_source.region_location
		replication_rule {
			rule_id = "%s"
			enable = true
			prefix = "%s"
			priority = 1
			deletemarker_replication_status = true
			destination_bucket_crn = ibm_cos_bucket.cos_bucket_destination.crn
		}
	}

	`, cosServiceNameSrc, bucketNameSrc, regionSrc, storageClassSrc, cosServiceNameDest, bucketNameDest, regionDest, storageClassDest, accountID, accountID, ruleId, prefix)
}
