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
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func TestAccIBMCmOfferingInstanceBasic(t *testing.T) {
	var conf catalogmanagementv1.OfferingInstance
	xAuthRefreshToken := fmt.Sprintf("tf_x_auth_refresh_token_%d", acctest.RandIntRange(10, 100))
	xAuthRefreshTokenUpdate := fmt.Sprintf("tf_x_auth_refresh_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmOfferingInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingInstanceConfigBasic(xAuthRefreshToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmOfferingInstanceExists("ibm_cm_offering_instance.cm_offering_instance", conf),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "x_auth_refresh_token", xAuthRefreshToken),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingInstanceConfigBasic(xAuthRefreshTokenUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "x_auth_refresh_token", xAuthRefreshTokenUpdate),
				),
			},
		},
	})
}

func TestAccIBMCmOfferingInstanceAllArgs(t *testing.T) {
	var conf catalogmanagementv1.OfferingInstance
	xAuthRefreshToken := fmt.Sprintf("tf_x_auth_refresh_token_%d", acctest.RandIntRange(10, 100))
	rev := fmt.Sprintf("tf_rev_%d", acctest.RandIntRange(10, 100))
	url := fmt.Sprintf("tf_url_%d", acctest.RandIntRange(10, 100))
	crn := fmt.Sprintf("tf_crn_%d", acctest.RandIntRange(10, 100))
	label := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	catalogID := fmt.Sprintf("tf_catalog_id_%d", acctest.RandIntRange(10, 100))
	offeringID := fmt.Sprintf("tf_offering_id_%d", acctest.RandIntRange(10, 100))
	kindFormat := fmt.Sprintf("tf_kind_format_%d", acctest.RandIntRange(10, 100))
	version := fmt.Sprintf("tf_version_%d", acctest.RandIntRange(10, 100))
	versionID := fmt.Sprintf("tf_version_id_%d", acctest.RandIntRange(10, 100))
	clusterID := fmt.Sprintf("tf_cluster_id_%d", acctest.RandIntRange(10, 100))
	clusterRegion := fmt.Sprintf("tf_cluster_region_%d", acctest.RandIntRange(10, 100))
	clusterAllNamespaces := "true"
	schematicsWorkspaceID := fmt.Sprintf("tf_schematics_workspace_id_%d", acctest.RandIntRange(10, 100))
	installPlan := fmt.Sprintf("tf_install_plan_%d", acctest.RandIntRange(10, 100))
	channel := fmt.Sprintf("tf_channel_%d", acctest.RandIntRange(10, 100))
	resourceGroupID := fmt.Sprintf("tf_resource_group_id_%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	disabled := "true"
	account := fmt.Sprintf("tf_account_%d", acctest.RandIntRange(10, 100))
	kindTarget := fmt.Sprintf("tf_kind_target_%d", acctest.RandIntRange(10, 100))
	sha := fmt.Sprintf("tf_sha_%d", acctest.RandIntRange(10, 100))
	xAuthRefreshTokenUpdate := fmt.Sprintf("tf_x_auth_refresh_token_%d", acctest.RandIntRange(10, 100))
	revUpdate := fmt.Sprintf("tf_rev_%d", acctest.RandIntRange(10, 100))
	urlUpdate := fmt.Sprintf("tf_url_%d", acctest.RandIntRange(10, 100))
	crnUpdate := fmt.Sprintf("tf_crn_%d", acctest.RandIntRange(10, 100))
	labelUpdate := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	catalogIDUpdate := fmt.Sprintf("tf_catalog_id_%d", acctest.RandIntRange(10, 100))
	offeringIDUpdate := fmt.Sprintf("tf_offering_id_%d", acctest.RandIntRange(10, 100))
	kindFormatUpdate := fmt.Sprintf("tf_kind_format_%d", acctest.RandIntRange(10, 100))
	versionUpdate := fmt.Sprintf("tf_version_%d", acctest.RandIntRange(10, 100))
	versionIDUpdate := fmt.Sprintf("tf_version_id_%d", acctest.RandIntRange(10, 100))
	clusterIDUpdate := fmt.Sprintf("tf_cluster_id_%d", acctest.RandIntRange(10, 100))
	clusterRegionUpdate := fmt.Sprintf("tf_cluster_region_%d", acctest.RandIntRange(10, 100))
	clusterAllNamespacesUpdate := "false"
	schematicsWorkspaceIDUpdate := fmt.Sprintf("tf_schematics_workspace_id_%d", acctest.RandIntRange(10, 100))
	installPlanUpdate := fmt.Sprintf("tf_install_plan_%d", acctest.RandIntRange(10, 100))
	channelUpdate := fmt.Sprintf("tf_channel_%d", acctest.RandIntRange(10, 100))
	resourceGroupIDUpdate := fmt.Sprintf("tf_resource_group_id_%d", acctest.RandIntRange(10, 100))
	locationUpdate := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	disabledUpdate := "false"
	accountUpdate := fmt.Sprintf("tf_account_%d", acctest.RandIntRange(10, 100))
	kindTargetUpdate := fmt.Sprintf("tf_kind_target_%d", acctest.RandIntRange(10, 100))
	shaUpdate := fmt.Sprintf("tf_sha_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmOfferingInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingInstanceConfig(xAuthRefreshToken, rev, url, crn, label, catalogID, offeringID, kindFormat, version, versionID, clusterID, clusterRegion, clusterAllNamespaces, schematicsWorkspaceID, installPlan, channel, resourceGroupID, location, disabled, account, kindTarget, sha),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmOfferingInstanceExists("ibm_cm_offering_instance.cm_offering_instance", conf),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "x_auth_refresh_token", xAuthRefreshToken),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "rev", rev),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "url", url),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "crn", crn),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "label", label),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "catalog_id", catalogID),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "offering_id", offeringID),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "kind_format", kindFormat),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "version", version),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "version_id", versionID),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "cluster_id", clusterID),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "cluster_region", clusterRegion),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "cluster_all_namespaces", clusterAllNamespaces),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "schematics_workspace_id", schematicsWorkspaceID),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "install_plan", installPlan),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "channel", channel),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "resource_group_id", resourceGroupID),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "location", location),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "disabled", disabled),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "account", account),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "kind_target", kindTarget),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "sha", sha),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingInstanceConfig(xAuthRefreshTokenUpdate, revUpdate, urlUpdate, crnUpdate, labelUpdate, catalogIDUpdate, offeringIDUpdate, kindFormatUpdate, versionUpdate, versionIDUpdate, clusterIDUpdate, clusterRegionUpdate, clusterAllNamespacesUpdate, schematicsWorkspaceIDUpdate, installPlanUpdate, channelUpdate, resourceGroupIDUpdate, locationUpdate, disabledUpdate, accountUpdate, kindTargetUpdate, shaUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "x_auth_refresh_token", xAuthRefreshTokenUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "rev", revUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "url", urlUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "crn", crnUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "label", labelUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "catalog_id", catalogIDUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "offering_id", offeringIDUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "kind_format", kindFormatUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "version", versionUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "version_id", versionIDUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "cluster_id", clusterIDUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "cluster_region", clusterRegionUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "cluster_all_namespaces", clusterAllNamespacesUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "schematics_workspace_id", schematicsWorkspaceIDUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "install_plan", installPlanUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "channel", channelUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "resource_group_id", resourceGroupIDUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "location", locationUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "disabled", disabledUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "account", accountUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "kind_target", kindTargetUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "sha", shaUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cm_offering_instance.cm_offering_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCmOfferingInstanceConfigBasic(xAuthRefreshToken string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_offering_instance" "cm_offering_instance" {
			x_auth_refresh_token = "%s"
		}
	`, xAuthRefreshToken)
}

func testAccCheckIBMCmOfferingInstanceConfig(xAuthRefreshToken string, rev string, url string, crn string, label string, catalogID string, offeringID string, kindFormat string, version string, versionID string, clusterID string, clusterRegion string, clusterAllNamespaces string, schematicsWorkspaceID string, installPlan string, channel string, resourceGroupID string, location string, disabled string, account string, kindTarget string, sha string) string {
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
	`, xAuthRefreshToken, rev, url, crn, label, catalogID, offeringID, kindFormat, version, versionID, clusterID, clusterRegion, clusterAllNamespaces, schematicsWorkspaceID, installPlan, channel, resourceGroupID, location, disabled, account, kindTarget, sha)
}

func testAccCheckIBMCmOfferingInstanceExists(n string, obj catalogmanagementv1.OfferingInstance) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getOfferingInstanceOptions := &catalogmanagementv1.GetOfferingInstanceOptions{}

		getOfferingInstanceOptions.SetInstanceIdentifier(rs.Primary.ID)

		offeringInstance, _, err := catalogManagementClient.GetOfferingInstance(getOfferingInstanceOptions)
		if err != nil {
			return err
		}

		obj = *offeringInstance
		return nil
	}
}

func testAccCheckIBMCmOfferingInstanceDestroy(s *terraform.State) error {
	catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cm_offering_instance" {
			continue
		}

		getOfferingInstanceOptions := &catalogmanagementv1.GetOfferingInstanceOptions{}

		getOfferingInstanceOptions.SetInstanceIdentifier(rs.Primary.ID)

		// Try to find the key
		_, response, err := catalogManagementClient.GetOfferingInstance(getOfferingInstanceOptions)

		if err == nil {
			return fmt.Errorf("cm_offering_instance still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cm_offering_instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
