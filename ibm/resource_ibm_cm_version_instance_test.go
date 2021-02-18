/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"fmt"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccIbmCmVersionInstanceBasic(t *testing.T) {
	var conf catalogmanagementv1.VersionInstance
	xAuthRefreshToken := fmt.Sprintf("X-Auth-Refresh-Token_%d", acctest.RandIntRange(10, 100))
	xAuthRefreshTokenUpdate := fmt.Sprintf("X-Auth-Refresh-Token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmCmVersionInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmCmVersionInstanceConfigBasic(xAuthRefreshToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCmVersionInstanceExists("ibm_cm_version_instance.cm_version_instance", conf),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "X-Auth-Refresh-Token", xAuthRefreshToken),
				),
			},
			{
				Config: testAccCheckIbmCmVersionInstanceConfigBasic(xAuthRefreshTokenUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "X-Auth-Refresh-Token", xAuthRefreshTokenUpdate),
				),
			},
		},
	})
}

func TestAccIbmCmVersionInstanceAllArgs(t *testing.T) {
	var conf catalogmanagementv1.VersionInstance
	xAuthRefreshToken := fmt.Sprintf("X-Auth-Refresh-Token_%d", acctest.RandIntRange(10, 100))
	id := fmt.Sprintf("id_%d", acctest.RandIntRange(10, 100))
	url := fmt.Sprintf("url_%d", acctest.RandIntRange(10, 100))
	crn := fmt.Sprintf("crn_%d", acctest.RandIntRange(10, 100))
	label := fmt.Sprintf("label_%d", acctest.RandIntRange(10, 100))
	catalogID := fmt.Sprintf("catalog_id_%d", acctest.RandIntRange(10, 100))
	offeringID := fmt.Sprintf("offering_id_%d", acctest.RandIntRange(10, 100))
	kindFormat := fmt.Sprintf("kind_format_%d", acctest.RandIntRange(10, 100))
	version := fmt.Sprintf("version_%d", acctest.RandIntRange(10, 100))
	clusterID := fmt.Sprintf("cluster_id_%d", acctest.RandIntRange(10, 100))
	clusterRegion := fmt.Sprintf("cluster_region_%d", acctest.RandIntRange(10, 100))
	clusterAllNamespaces := "true"
	xAuthRefreshTokenUpdate := fmt.Sprintf("X-Auth-Refresh-Token_%d", acctest.RandIntRange(10, 100))
	idUpdate := fmt.Sprintf("id_%d", acctest.RandIntRange(10, 100))
	urlUpdate := fmt.Sprintf("url_%d", acctest.RandIntRange(10, 100))
	crnUpdate := fmt.Sprintf("crn_%d", acctest.RandIntRange(10, 100))
	labelUpdate := fmt.Sprintf("label_%d", acctest.RandIntRange(10, 100))
	catalogIDUpdate := fmt.Sprintf("catalog_id_%d", acctest.RandIntRange(10, 100))
	offeringIDUpdate := fmt.Sprintf("offering_id_%d", acctest.RandIntRange(10, 100))
	kindFormatUpdate := fmt.Sprintf("kind_format_%d", acctest.RandIntRange(10, 100))
	versionUpdate := fmt.Sprintf("version_%d", acctest.RandIntRange(10, 100))
	clusterIDUpdate := fmt.Sprintf("cluster_id_%d", acctest.RandIntRange(10, 100))
	clusterRegionUpdate := fmt.Sprintf("cluster_region_%d", acctest.RandIntRange(10, 100))
	clusterAllNamespacesUpdate := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmCmVersionInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmCmVersionInstanceConfig(xAuthRefreshToken, id, url, crn, label, catalogID, offeringID, kindFormat, version, clusterID, clusterRegion, clusterAllNamespaces),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCmVersionInstanceExists("ibm_cm_version_instance.cm_version_instance", conf),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "X-Auth-Refresh-Token", xAuthRefreshToken),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "id", id),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "url", url),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "crn", crn),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "label", label),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "catalog_id", catalogID),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "offering_id", offeringID),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "kind_format", kindFormat),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "version", version),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "cluster_id", clusterID),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "cluster_region", clusterRegion),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "cluster_all_namespaces", clusterAllNamespaces),
				),
			},
			{
				Config: testAccCheckIbmCmVersionInstanceConfig(xAuthRefreshTokenUpdate, idUpdate, urlUpdate, crnUpdate, labelUpdate, catalogIDUpdate, offeringIDUpdate, kindFormatUpdate, versionUpdate, clusterIDUpdate, clusterRegionUpdate, clusterAllNamespacesUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "X-Auth-Refresh-Token", xAuthRefreshTokenUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "id", idUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "url", urlUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "crn", crnUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "label", labelUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "catalog_id", catalogIDUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "offering_id", offeringIDUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "kind_format", kindFormatUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "version", versionUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "cluster_id", clusterIDUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "cluster_region", clusterRegionUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version_instance.cm_version_instance", "cluster_all_namespaces", clusterAllNamespacesUpdate),
				),
			},
			{
				ResourceName:      "ibm_cm_version_instance.cm_version_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmCmVersionInstanceConfigBasic(xAuthRefreshToken string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_version_instance" "cm_version_instance" {
			X-Auth-Refresh-Token = "%s"
		}
	`, xAuthRefreshToken)
}

func testAccCheckIbmCmVersionInstanceConfig(xAuthRefreshToken string, id string, url string, crn string, label string, catalogID string, offeringID string, kindFormat string, version string, clusterID string, clusterRegion string, clusterAllNamespaces string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_version_instance" "cm_version_instance" {
			X-Auth-Refresh-Token = "%s"
			id = "%s"
			url = "%s"
			crn = "%s"
			label = "%s"
			catalog_id = "%s"
			offering_id = "%s"
			kind_format = "%s"
			version = "%s"
			cluster_id = "%s"
			cluster_region = "%s"
			cluster_namespaces = "FIXME"
			cluster_all_namespaces = %s
		}
	`, xAuthRefreshToken, id, url, crn, label, catalogID, offeringID, kindFormat, version, clusterID, clusterRegion, clusterAllNamespaces)
}

func testAccCheckIbmCmVersionInstanceExists(n string, obj catalogmanagementv1.VersionInstance) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getVersionInstanceOptions := &catalogmanagementv1.GetVersionInstanceOptions{}

		getVersionInstanceOptions.SetInstanceIdentifier(rs.Primary.ID)

		versionInstance, _, err := catalogManagementClient.GetVersionInstance(getVersionInstanceOptions)
		if err != nil {
			return err
		}

		obj = *versionInstance
		return nil
	}
}

func testAccCheckIbmCmVersionInstanceDestroy(s *terraform.State) error {
	catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cm_version_instance" {
			continue
		}

		getVersionInstanceOptions := &catalogmanagementv1.GetVersionInstanceOptions{}

		getVersionInstanceOptions.SetInstanceIdentifier(rs.Primary.ID)

		// Try to find the key
		_, response, err := catalogManagementClient.GetVersionInstance(getVersionInstanceOptions)

		if err == nil {
			return fmt.Errorf("A version instance resource (provision instance of a catalog version). still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for A version instance resource (provision instance of a catalog version). (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
