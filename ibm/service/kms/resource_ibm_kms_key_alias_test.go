package kms_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMKMSResource_Key_Alias_Name(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	// cosInstanceName := fmt.Sprintf("cos_%d", acctest.RandIntRange(10, 100))
	// bucketName := fmt.Sprintf("bucket-test77")
	aliasName := fmt.Sprintf("alias_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsResourceAliasConfig(instanceName, keyName, aliasName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key_alias.testAlias", "alias", aliasName),
					resource.TestCheckResourceAttr("data.ibm_kms_keys.AliasTest", "alias", aliasName),
				),
			},
		},
	})
}
func TestAccIBMKMSResource_Key_Alias_Key(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	// cosInstanceName := fmt.Sprintf("cos_%d", acctest.RandIntRange(10, 100))
	// bucketName := fmt.Sprintf("bucket-test77")
	aliasName := fmt.Sprintf("alias_%d", acctest.RandIntRange(10, 100))

	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsResourceAliasDuplicateConfig(instanceName, keyName, aliasName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key_alias.testAlias", "alias", aliasName),
				),
			},
		},
	})
}

func TestAccIBMKMSResource_Key_Alias_Key_Duplicacy(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	// cosInstanceName := fmt.Sprintf("cos_%d", acctest.RandIntRange(10, 100))
	// bucketName := fmt.Sprintf("bucket-test77")
	aliasName := fmt.Sprintf("alias_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsResourceAliasDuplicateConfig(instanceName, keyName, aliasName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key_alias.testAlias", "alias", aliasName),
				),
			},
		},
	})
}

func TestAccIBMKMSResource_Key_Alias_Key_Check(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	// cosInstanceName := fmt.Sprintf("cos_%d", acctest.RandIntRange(10, 100))
	// bucketName := fmt.Sprintf("bucket-test77")
	aliasName := fmt.Sprintf("alias_%d", acctest.RandIntRange(10, 100))
	aliasName2 := fmt.Sprintf("alias_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	keyName2 := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsResourceAliasTwo(instanceName, keyName, aliasName, aliasName2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key_alias.testAlias", "alias", aliasName),
				),
			},
			{
				Config: testAccCheckIBMKmsResourceAliasOne(instanceName, keyName, aliasName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key_alias.testAlias", "alias", aliasName),
				),
			},
			{
				Config: testAccCheckIBMKmsResourceAliasOne(instanceName, keyName2, aliasName2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName2),
					resource.TestCheckResourceAttr("ibm_kms_key_alias.testAlias", "alias", aliasName2),
				),
			},
			{
				Config: testAccCheckIBMKmsResourceAliasWithExistingAlias(instanceName, keyName, aliasName, aliasName2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key_alias.testAlias", "alias", aliasName),
					resource.TestCheckResourceAttr("ibm_kms_key_alias.testAlias2", "existing_alias", aliasName),
					resource.TestCheckResourceAttr("ibm_kms_key_alias.testAlias2", "alias", aliasName2),
				),
			},
		},
	})
}

func TestAccIBMKMSResource_Key_Alias_Key_Limit(t *testing.T) {
	instanceName := fmt.Sprintf("tf_kms_%d", acctest.RandIntRange(10, 100))
	// cosInstanceName := fmt.Sprintf("cos_%d", acctest.RandIntRange(10, 100))
	// bucketName := fmt.Sprintf("bucket-test77")
	aliasName := fmt.Sprintf("alias_%d", acctest.RandIntRange(10, 100))
	aliasName2 := fmt.Sprintf("alias_%d", acctest.RandIntRange(10, 100))
	aliasName3 := fmt.Sprintf("alias_%d", acctest.RandIntRange(10, 100))
	aliasName4 := fmt.Sprintf("alias_%d", acctest.RandIntRange(10, 100))
	aliasName5 := fmt.Sprintf("alias_%d", acctest.RandIntRange(10, 100))
	aliasName6 := fmt.Sprintf("alias_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMKmsResourceAliasLimitConfig(instanceName, keyName, aliasName, aliasName2, aliasName3, aliasName4, aliasName5, aliasName6),
				ExpectError: regexp.MustCompile("(KEY_ALIAS_QUOTA_ERR)"),
			},
		},
	})
}

func testAccCheckIBMKmsResourceAliasConfig(instanceName, KeyName, aliasName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "%s"
		standard_key =  true
		force_delete = true
	}
	resource "ibm_kms_key_alias" "testAlias" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		alias = "%s"
		key_id = "${ibm_kms_key.test.key_id}"
	}
	data "ibm_kms_keys" "AliasTest" {
		instance_id = ibm_kms_key_alias.testAlias.instance_id
		alias = "${ibm_kms_key_alias.testAlias.alias}"
	}
`, instanceName, KeyName, aliasName)
}

func testAccCheckIBMKmsResourceAliasDuplicateConfig(instanceName, KeyName, aliasName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "%s"
		standard_key =  true
		force_delete = true
	}
	resource "ibm_kms_key_alias" "testAlias" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		alias = "%s"
		key_id = "${ibm_kms_key.test.key_id}"
	}
	resource "ibm_kms_key_alias" "testAlias2" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		alias = "${ibm_kms_key_alias.testAlias2.alias}"
		key_id = "${ibm_kms_key.test.key_id}"
	}

`, instanceName, KeyName, aliasName)
}

func testAccCheckIBMKmsResourceAliasTwo(instanceName, KeyName, aliasName, aliasName2 string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "%s"
		standard_key =  true
		force_delete = true
	}
	resource "ibm_kms_key_alias" "testAlias" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		alias = "%s"
		key_id = "${ibm_kms_key.test.key_id}"
	}
	resource "ibm_kms_key_alias" "testAlias2" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		alias = "%s"
		key_id = "${ibm_kms_key.test.key_id}"
	}

`, instanceName, KeyName, aliasName, aliasName2)
}

func testAccCheckIBMKmsResourceAliasWithExistingAlias(instanceName, KeyName, aliasName, aliasName2 string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "%s"
		standard_key =  true
		force_delete = true
	}
	resource "ibm_kms_key_alias" "testAlias" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		alias = "%s"
		key_id = "${ibm_kms_key.test.key_id}"
	}
	resource "ibm_kms_key_alias" "testAlias2" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		alias = "%s"
		existing_alias = "${ibm_kms_key_alias.testAlais.alias}"
	}

`, instanceName, KeyName, aliasName, aliasName2)
}

func testAccCheckIBMKmsResourceAliasOne(instanceName, KeyName, aliasName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "%s"
		standard_key =  true
		force_delete = true
	}
	resource "ibm_kms_key_alias" "testAlias" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		alias = "%s"
		key_id = "${ibm_kms_key.test.key_id}"
	}

`, instanceName, KeyName, aliasName)
}

func testAccCheckIBMKmsResourceAliasLimitConfig(instanceName, KeyName, aliasName, aliasName2, aliasName3, aliasName4, aliasName5, aliasName6 string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "%s"
		standard_key =  true
		force_delete = true
	}
	resource "ibm_kms_key_alias" "testAlias" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		alias = "%s"
		key_id = "${ibm_kms_key.test.key_id}"
	}
	resource "ibm_kms_key_alias" "testAlias2" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		alias = "%s"
		key_id = "${ibm_kms_key.test.key_id}"
	}
	resource "ibm_kms_key_alias" "testAlias3" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		alias = "%s"
		key_id = "${ibm_kms_key.test.key_id}"
	}
	resource "ibm_kms_key_alias" "testAlias4" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		alias = "%s"
		key_id = "${ibm_kms_key.test.key_id}"
	}
	resource "ibm_kms_key_alias" "testAlias5" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		alias = "%s"
		key_id = "${ibm_kms_key.test.key_id}"
	}
	resource "ibm_kms_key_alias" "testAlias6" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		alias = "%s"
		key_id = "${ibm_kms_key.test.key_id}"
	}
`, instanceName, KeyName, aliasName, aliasName2, aliasName3, aliasName4, aliasName5, aliasName6)
}
