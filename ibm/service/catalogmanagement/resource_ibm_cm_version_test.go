// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func TestAccIBMCmVersionBasic(t *testing.T) {
	var conf catalogmanagementv1.Version

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmVersionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmVersionConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmVersionExists("ibm_cm_version.cm_version", conf),
				),
			},
		},
	})
}

func TestAccIBMCmVersionSimpleArgs(t *testing.T) {
	var conf catalogmanagementv1.Version
	zipurl := "https://github.com/IBM-Cloud/terraform-sample/archive/refs/tags/v1.1.0.tar.gz"
	targetVersion := "2.2.2"
	includeConfig := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmVersionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmVersionSimpleConfig(zipurl, targetVersion, includeConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmVersionExists("ibm_cm_version.cm_version", conf),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "zipurl", zipurl),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "target_version", targetVersion),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "include_config", includeConfig),
				),
			},
		},
	})
}

func TestAccIBMCmVersionComplexArgs(t *testing.T) {
	var conf catalogmanagementv1.Version
	zipurl := "https://github.com/IBM-Cloud/terraform-sample/archive/refs/tags/v1.1.0.tar.gz"
	targetVersion := "1.0.0"
	usageText := "example usage text"
	installInstructions := "example install instructions"
	iamPermissionServiceNameOne := "iam-identity"
	iamPermissionsRoleCRNsOne := "[\"crn:v1:bluemix:public:iam::::role:Administrator\"]"
	iamPermissionServiceNameTwo := "is.vpc"
	iamPermissionsRoleCRNsTwo := "[\"crn:v1:bluemix:public:iam::::serviceRole:Manager\", \"crn:v1:bluemix:public:iam::::role:Administrator\"]"
	featureTitleOne := "RAM"
	featureDescriptionOne := "8GB"
	featureTitleTwo := "Storage"
	featureDescriptionTwo := "500GB"
	archDiagDesc := "diagram 1 description"
	archDiagCaption := "diagram 1 caption"
	archDiagType := "image/svg+xml"
	archDiagURL := "https://raw.githubusercontent.com/gmendel/landing-zone-vpc/main/.docs/vpc-standard.svg"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmVersionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmVersionComplexConfig(zipurl, targetVersion, usageText, installInstructions, iamPermissionServiceNameOne, iamPermissionsRoleCRNsOne, iamPermissionServiceNameTwo, iamPermissionsRoleCRNsTwo, featureTitleOne, featureDescriptionOne, featureTitleTwo, featureDescriptionTwo, archDiagDesc, archDiagCaption, archDiagType, archDiagURL),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmVersionExists("ibm_cm_version.cm_version", conf),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "zipurl", zipurl),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "target_version", targetVersion),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "usage", usageText),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "install.0.instructions", installInstructions),
					resource.TestCheckResourceAttrSet("ibm_cm_version.cm_version", "iam_permissions.0.service_name"),
					resource.TestCheckResourceAttrSet("ibm_cm_version.cm_version", "iam_permissions.0.role_crns.0"),
					resource.TestCheckResourceAttrSet("ibm_cm_version.cm_version", "iam_permissions.1.service_name"),
					resource.TestCheckResourceAttrSet("ibm_cm_version.cm_version", "iam_permissions.1.role_crns.0"),
					resource.TestCheckResourceAttrSet("ibm_cm_version.cm_version", "flavor.#"),
					resource.TestCheckResourceAttrSet("ibm_cm_version.cm_version", "solution_info.#"),
				),
			},
		},
	})
}

func TestAccIBMCmVersionVSI(t *testing.T) {
	var conf catalogmanagementv1.Version
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	label := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	installKind := "instance"
	sha := "64245e5f3f1e9c4048b18db3abd1450d4b6f9e263ac1b33df6fc1ae96fcbdebb"
	targetVersion := "3.3.3"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmVersionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmVersionVSIConfig(name, label, installKind, sha, targetVersion),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmVersionExists("ibm_cm_version.cm_version", conf),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "name", name),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "label", label),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "install_kind", installKind),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "sha", sha),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "target_version", targetVersion),
				),
			},
		},
	})
}

func testAccCheckIBMCmVersionConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
			label = "test_tf_catalog_label_1"
			kind = "offering"
		}

		resource "ibm_cm_offering" "cm_offering" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			label = "test_tf_offering_label_1"
			name = "test_tf_offering_name_1"
			offering_icon_url = "test.url.1"
			tags = ["dev_ops"]
		}

		resource "ibm_cm_version" "cm_version" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.id
			zipurl = "https://github.com/IBM-Cloud/terraform-sample/archive/refs/tags/v1.1.0.tar.gz"
			install {}
		}
	`)
}

func testAccCheckIBMCmVersionSimpleConfig(zipurl string, targetVersion string, includeConfig string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "test_tf_catalog_label_2"
			kind = "offering"
		}

		resource "ibm_cm_offering" "cm_offering" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			label = "test_tf_offering_label_2"
			name = "test_tf_offering_name_2"
			offering_icon_url = "test.url.2"
			tags = ["dev_ops"]
		}

		resource "ibm_cm_version" "cm_version" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.id
			zipurl = "%s"
			target_version = "%s"
			include_config = %s
			install {}
		}
	`, zipurl, targetVersion, includeConfig)
}

func testAccCheckIBMCmVersionComplexConfig(zipurl string, targetVersion string, usageText string, installInstructions string, iamPermissionServiceNameOne string, iamPermissionsRoleCRNsOne string, iamPermissionServiceNameTwo string, iamPermissionsRoleCRNsTwo string, featureTitleOne string, featureDescriptionOne string, featureTitleTwo string, featureDescriptionTwo string, archDiagDesc string, archDiagCaption string, archDiagType string, archDiagURL string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "test_tf_catalog_label_2"
			kind = "offering"
		}

		resource "ibm_cm_offering" "cm_offering" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			label = "test_tf_offering_label_2"
			name = "test_tf_offering_name_2"
			offering_icon_url = "test.url.2"
			tags = ["dev_ops"]
			product_kind = "solution"
		}

		resource "ibm_cm_version" "cm_version" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.id
			zipurl = "%s"
			target_version = "%s"
			include_config = true
			usage = "%s"
			flavor {
				name = "standard"
				label = "Standard"
				index = 1
			}
			configuration {
				default_value = "foo"
				description = "The name to pass to the template."
				key = "name"
				type = "string"
				hidden = false
				required = false
				value_constraints {
					type  = "regex"
					value = "*"
					description = "Invalid name input"
				}
			}
			install {
				instructions = "%s"
			}
			iam_permissions {
				service_name = "%s"
				role_crns = %s
			}
			iam_permissions {
				service_name = "%s"
				role_crns = %s
			}
			solution_info {
				features {
					title = "%s"
					description = "%s"
				}
				features {
					title = "%s"
					description = "%s"
				}
				architecture_diagrams {
					description = "%s"
					diagram {
						caption = "%s"
						type = "%s"
						url = "%s"
					}
				}
			}
			required_resources {
				type = "terraformVersion"
				value = "1.5"
			}
		}
	`, zipurl, targetVersion, usageText, installInstructions, iamPermissionServiceNameOne, iamPermissionsRoleCRNsOne, iamPermissionServiceNameTwo, iamPermissionsRoleCRNsTwo, featureTitleOne, featureDescriptionOne, featureTitleTwo, featureDescriptionTwo, archDiagDesc, archDiagCaption, archDiagType, archDiagURL)
}

func testAccCheckIBMCmVersionVSIConfig(name string, label string, installKind string, sha string, targetVersion string) string {
	return fmt.Sprintf(`

	resource "ibm_cm_catalog" "cm_catalog" {
		label = "test_tf_catalog_label_3"
		kind = "offering"
	}

	resource "ibm_cm_offering" "cm_offering" {
		catalog_id = ibm_cm_catalog.cm_catalog.id
		label = "test_tf_offering_label_3"
		name = "test_tf_offering_name_3"
		offering_icon_url = "test.url.2"
		tags = ["dev_ops"]
	}

		resource "ibm_cm_version" "cm_version" {
			name = "%s"
			label = "%s"
			catalog_id = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.id
			tags = ["virtualservers"]
			target_kinds = [ "vpc-x86" ]
			install_kind = "%s"
			import_sha = "%s"
			target_version = "%s"
			install {}

			import_metadata {
				operating_system {
					dedicated_host_only = false
					vendor = "CentOS"
					name = "centos-7-amd64"
					href = "https://us-south-stage01.iaasdev.cloud.ibm.com/v1/operating_systems/centos-7-amd64"
					display_name = "CentOS 7.x - Minimal Install (amd64)"
					family = "CentOS"
					version = "7.x - Minimal Install"
					architecture = "amd64"
				}
				minimum_provisioned_size = 100
				file {
					size = 1
				}
				images {
					id = "r134-7fafcc04-f09c-4959-bed5-f6b655409c7b"
					name = "dubee-test-2"
					region = "us-south"
				}
			}
		}
	`, name, label, installKind, sha, targetVersion)
}

func testAccCheckIBMCmVersionExists(n string, obj catalogmanagementv1.Version) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getVersionOptions := &catalogmanagementv1.GetVersionOptions{}
		getVersionOptions.SetVersionLocID(strings.Replace(rs.Primary.ID, "/", ".", 1))

		offering, _, err := catalogManagementClient.GetVersion(getVersionOptions)
		version := offering.Kinds[0].Versions[0]
		if err != nil {
			return err
		}

		obj = version
		return nil
	}
}

func testAccCheckIBMCmVersionDestroy(s *terraform.State) error {
	catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cm_version" {
			continue
		}

		getVersionOptions := &catalogmanagementv1.GetVersionOptions{}
		getVersionOptions.SetVersionLocID(strings.Replace(rs.Primary.ID, "/", ".", 1))

		// Try to find the key
		_, response, err := catalogManagementClient.GetVersion(getVersionOptions)

		if err == nil {
			return fmt.Errorf("cm_version still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cm_version (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
