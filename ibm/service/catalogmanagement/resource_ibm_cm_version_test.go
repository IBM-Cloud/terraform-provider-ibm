// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
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

func TestAccIBMCmVersionAllArgs(t *testing.T) {
	var conf catalogmanagementv1.Version
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	label := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	installKind := fmt.Sprintf("tf_install_kind_%d", acctest.RandIntRange(10, 100))
	formatKind := fmt.Sprintf("tf_format_kind_%d", acctest.RandIntRange(10, 100))
	productKind := fmt.Sprintf("tf_product_kind_%d", acctest.RandIntRange(10, 100))
	sha := fmt.Sprintf("tf_sha_%d", acctest.RandIntRange(10, 100))
	version := fmt.Sprintf("tf_version_%d", acctest.RandIntRange(10, 100))
	workingDirectory := fmt.Sprintf("tf_working_directory_%d", acctest.RandIntRange(10, 100))
	zipurl := fmt.Sprintf("tf_zipurl_%d", acctest.RandIntRange(10, 100))
	targetVersion := fmt.Sprintf("tf_target_version_%d", acctest.RandIntRange(10, 100))
	includeConfig := "true"
	isVsi := "true"
	repotype := fmt.Sprintf("tf_repotype_%d", acctest.RandIntRange(10, 100))
	xAuthToken := fmt.Sprintf("tf_x_auth_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmVersionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmVersionConfig(name, label, installKind, formatKind, productKind, sha, version, workingDirectory, zipurl, targetVersion, includeConfig, isVsi, repotype, xAuthToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmVersionExists("ibm_cm_version.cm_version", conf),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "name", name),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "label", label),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "install_kind", installKind),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "format_kind", formatKind),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "product_kind", productKind),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "sha", sha),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "version", version),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "working_directory", workingDirectory),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "zipurl", zipurl),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "target_version", targetVersion),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "include_config", includeConfig),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "is_vsi", isVsi),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "repotype", repotype),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "x_auth_token", xAuthToken),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cm_version.cm_version",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCmVersionConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
		}

		resource "ibm_cm_version" "cm_version" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.offering_id
		}
	`)
}

func testAccCheckIBMCmVersionConfig(name string, label string, installKind string, formatKind string, productKind string, sha string, version string, workingDirectory string, zipurl string, targetVersion string, includeConfig string, isVsi string, repotype string, xAuthToken string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
		}

		resource "ibm_cm_version" "cm_version" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.offering_id
			tags = "FIXME"
			content = "FIXME"
			name = "%s"
			label = "%s"
			install_kind = "%s"
			target_kinds = "FIXME"
			format_kind = "%s"
			product_kind = "%s"
			sha = "%s"
			version = "%s"
			flavor {
				name = "name"
				label = "label"
				label_i18n = { "key": "inner" }
				index = 1
			}
			metadata {
				operating_system {
					dedicated_host_only = true
					vendor = "vendor"
					name = "name"
					href = "href"
					display_name = "display_name"
					family = "family"
					version = "version"
					architecture = "architecture"
				}
				file {
					size = 1
				}
				minimum_provisioned_size = 1
				images {
					id = "id"
					name = "name"
					region = "region"
				}
			}
			working_directory = "%s"
			zipurl = "%s"
			target_version = "%s"
			include_config = %s
			is_vsi = %s
			repotype = "%s"
			x_auth_token = "%s"
		}
	`, name, label, installKind, formatKind, productKind, sha, version, workingDirectory, zipurl, targetVersion, includeConfig, isVsi, repotype, xAuthToken)
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

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getVersionOptions.SetCatalogIdentifier(parts[0])
		getVersionOptions.SetOfferingID(parts[1])
		getVersionOptions.SetVersionLocID(parts[2])

		version, _, err := catalogManagementClient.GetVersion(getVersionOptions)
		if err != nil {
			return err
		}

		obj = *version
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

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getVersionOptions.SetCatalogIdentifier(parts[0])
		getVersionOptions.SetOfferingID(parts[1])
		getVersionOptions.SetVersionLocID(parts[2])

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
