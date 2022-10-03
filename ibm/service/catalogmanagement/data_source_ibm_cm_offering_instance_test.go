// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCmOfferingInstanceDataSourceBasic(t *testing.T) {
	offeringInstanceXAuthRefreshToken := fmt.Sprintf("tf_x_auth_refresh_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingInstanceDataSourceConfigBasic(offeringInstanceXAuthRefreshToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "instance_identifier"),
				),
			},
		},
	})
}

func TestAccIBMCmOfferingInstanceDataSourceAllArgs(t *testing.T) {
	offeringInstanceXAuthRefreshToken := fmt.Sprintf("tf_x_auth_refresh_token_%d", acctest.RandIntRange(10, 100))
	offeringInstanceRev := fmt.Sprintf("tf_rev_%d", acctest.RandIntRange(10, 100))
	offeringInstanceURL := fmt.Sprintf("tf_url_%d", acctest.RandIntRange(10, 100))
	offeringInstanceCRN := fmt.Sprintf("tf_crn_%d", acctest.RandIntRange(10, 100))
	offeringInstanceLabel := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	offeringInstanceCatalogID := fmt.Sprintf("tf_catalog_id_%d", acctest.RandIntRange(10, 100))
	offeringInstanceOfferingID := fmt.Sprintf("tf_offering_id_%d", acctest.RandIntRange(10, 100))
	offeringInstanceKindFormat := fmt.Sprintf("tf_kind_format_%d", acctest.RandIntRange(10, 100))
	offeringInstanceVersion := fmt.Sprintf("tf_version_%d", acctest.RandIntRange(10, 100))
	offeringInstanceVersionID := fmt.Sprintf("tf_version_id_%d", acctest.RandIntRange(10, 100))
	offeringInstanceClusterID := fmt.Sprintf("tf_cluster_id_%d", acctest.RandIntRange(10, 100))
	offeringInstanceClusterRegion := fmt.Sprintf("tf_cluster_region_%d", acctest.RandIntRange(10, 100))
	offeringInstanceClusterAllNamespaces := "true"
	offeringInstanceSchematicsWorkspaceID := fmt.Sprintf("tf_schematics_workspace_id_%d", acctest.RandIntRange(10, 100))
	offeringInstanceInstallPlan := fmt.Sprintf("tf_install_plan_%d", acctest.RandIntRange(10, 100))
	offeringInstanceChannel := fmt.Sprintf("tf_channel_%d", acctest.RandIntRange(10, 100))
	offeringInstanceResourceGroupID := fmt.Sprintf("tf_resource_group_id_%d", acctest.RandIntRange(10, 100))
	offeringInstanceLocation := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	offeringInstanceDisabled := "true"
	offeringInstanceAccount := fmt.Sprintf("tf_account_%d", acctest.RandIntRange(10, 100))
	offeringInstanceKindTarget := fmt.Sprintf("tf_kind_target_%d", acctest.RandIntRange(10, 100))
	offeringInstanceSha := fmt.Sprintf("tf_sha_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingInstanceDataSourceConfig(offeringInstanceXAuthRefreshToken, offeringInstanceRev, offeringInstanceURL, offeringInstanceCRN, offeringInstanceLabel, offeringInstanceCatalogID, offeringInstanceOfferingID, offeringInstanceKindFormat, offeringInstanceVersion, offeringInstanceVersionID, offeringInstanceClusterID, offeringInstanceClusterRegion, offeringInstanceClusterAllNamespaces, offeringInstanceSchematicsWorkspaceID, offeringInstanceInstallPlan, offeringInstanceChannel, offeringInstanceResourceGroupID, offeringInstanceLocation, offeringInstanceDisabled, offeringInstanceAccount, offeringInstanceKindTarget, offeringInstanceSha),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "instance_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "offering_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "kind_format"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "version_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "cluster_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "cluster_region"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "cluster_namespaces.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "cluster_all_namespaces"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "schematics_workspace_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "install_plan"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "channel"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "metadata.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "disabled"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "account"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "last_operation.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "kind_target"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "sha"),
				),
			},
		},
	})
}

func testAccCheckIBMCmOfferingInstanceDataSourceConfigBasic(offeringInstanceXAuthRefreshToken string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_offering_instance" "cm_offering_instance" {
			x_auth_refresh_token = "%s"
		}

		data "ibm_cm_offering_instance" "cm_offering_instance" {
			instance_identifier = "instance_identifier"
		}
	`, offeringInstanceXAuthRefreshToken)
}

func testAccCheckIBMCmOfferingInstanceDataSourceConfig(offeringInstanceXAuthRefreshToken string, offeringInstanceRev string, offeringInstanceURL string, offeringInstanceCRN string, offeringInstanceLabel string, offeringInstanceCatalogID string, offeringInstanceOfferingID string, offeringInstanceKindFormat string, offeringInstanceVersion string, offeringInstanceVersionID string, offeringInstanceClusterID string, offeringInstanceClusterRegion string, offeringInstanceClusterAllNamespaces string, offeringInstanceSchematicsWorkspaceID string, offeringInstanceInstallPlan string, offeringInstanceChannel string, offeringInstanceResourceGroupID string, offeringInstanceLocation string, offeringInstanceDisabled string, offeringInstanceAccount string, offeringInstanceKindTarget string, offeringInstanceSha string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_offering_instance" "cm_offering_instance" {
			x_auth_refresh_token = "%s"
			rev = "%s"
			url = "%s"
			crn = "%s"
			label = "%s"
			catalog_id = "%s"
			offering_id = "%s"
			kind_format = "%s"
			version = "%s"
			version_id = "%s"
			cluster_id = "%s"
			cluster_region = "%s"
			cluster_namespaces = "FIXME"
			cluster_all_namespaces = %s
			schematics_workspace_id = "%s"
			install_plan = "%s"
			channel = "%s"
			created = "2004-10-28T04:39:00.000Z"
			updated = "2004-10-28T04:39:00.000Z"
			metadata = "FIXME"
			resource_group_id = "%s"
			location = "%s"
			disabled = %s
			account = "%s"
			last_operation {
				operation = "operation"
				state = "state"
				message = "message"
				transaction_id = "transaction_id"
				updated = "2021-01-31T09:44:12Z"
				code = "code"
			}
			kind_target = "%s"
			sha = "%s"
		}

		data "ibm_cm_offering_instance" "cm_offering_instance" {
			instance_identifier = "instance_identifier"
		}
	`, offeringInstanceXAuthRefreshToken, offeringInstanceRev, offeringInstanceURL, offeringInstanceCRN, offeringInstanceLabel, offeringInstanceCatalogID, offeringInstanceOfferingID, offeringInstanceKindFormat, offeringInstanceVersion, offeringInstanceVersionID, offeringInstanceClusterID, offeringInstanceClusterRegion, offeringInstanceClusterAllNamespaces, offeringInstanceSchematicsWorkspaceID, offeringInstanceInstallPlan, offeringInstanceChannel, offeringInstanceResourceGroupID, offeringInstanceLocation, offeringInstanceDisabled, offeringInstanceAccount, offeringInstanceKindTarget, offeringInstanceSha)
}
