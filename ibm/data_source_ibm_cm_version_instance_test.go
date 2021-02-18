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
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccIbmCmVersionInstanceDataSourceBasic(t *testing.T) {
	versionInstanceXAuthRefreshToken := fmt.Sprintf("x_auth_refresh_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmCmVersionInstanceDataSourceConfigBasic(versionInstanceXAuthRefreshToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "instance_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "offering_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "kind_format"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "cluster_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "cluster_region"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "cluster_namespaces.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "cluster_all_namespaces"),
				),
			},
		},
	})
}

func TestAccIbmCmVersionInstanceDataSourceAllArgs(t *testing.T) {
	versionInstanceXAuthRefreshToken := fmt.Sprintf("x_auth_refresh_token_%d", acctest.RandIntRange(10, 100))
	versionInstanceID := fmt.Sprintf("id_%d", acctest.RandIntRange(10, 100))
	versionInstanceURL := fmt.Sprintf("url_%d", acctest.RandIntRange(10, 100))
	versionInstanceCrn := fmt.Sprintf("crn_%d", acctest.RandIntRange(10, 100))
	versionInstanceLabel := fmt.Sprintf("label_%d", acctest.RandIntRange(10, 100))
	versionInstanceCatalogID := fmt.Sprintf("catalog_id_%d", acctest.RandIntRange(10, 100))
	versionInstanceOfferingID := fmt.Sprintf("offering_id_%d", acctest.RandIntRange(10, 100))
	versionInstanceKindFormat := fmt.Sprintf("kind_format_%d", acctest.RandIntRange(10, 100))
	versionInstanceVersion := fmt.Sprintf("version_%d", acctest.RandIntRange(10, 100))
	versionInstanceClusterID := fmt.Sprintf("cluster_id_%d", acctest.RandIntRange(10, 100))
	versionInstanceClusterRegion := fmt.Sprintf("cluster_region_%d", acctest.RandIntRange(10, 100))
	versionInstanceClusterAllNamespaces := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmCmVersionInstanceDataSourceConfig(versionInstanceXAuthRefreshToken, versionInstanceID, versionInstanceURL, versionInstanceCrn, versionInstanceLabel, versionInstanceCatalogID, versionInstanceOfferingID, versionInstanceKindFormat, versionInstanceVersion, versionInstanceClusterID, versionInstanceClusterRegion, versionInstanceClusterAllNamespaces),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "instance_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "offering_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "kind_format"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "cluster_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "cluster_region"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "cluster_namespaces.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version_instance.cm_version_instance", "cluster_all_namespaces"),
				),
			},
		},
	})
}

func testAccCheckIbmCmVersionInstanceDataSourceConfigBasic(versionInstanceXAuthRefreshToken string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_version_instance" "cm_version_instance" {
			X-Auth-Refresh-Token = "%s"
		}

		data "ibm_cm_version_instance" "cm_version_instance" {
			instance_identifier = "instance_identifier"
		}
	`, versionInstanceXAuthRefreshToken)
}

func testAccCheckIbmCmVersionInstanceDataSourceConfig(versionInstanceXAuthRefreshToken string, versionInstanceID string, versionInstanceURL string, versionInstanceCrn string, versionInstanceLabel string, versionInstanceCatalogID string, versionInstanceOfferingID string, versionInstanceKindFormat string, versionInstanceVersion string, versionInstanceClusterID string, versionInstanceClusterRegion string, versionInstanceClusterAllNamespaces string) string {
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

		data "ibm_cm_version_instance" "cm_version_instance" {
			instance_identifier = "instance_identifier"
		}
	`, versionInstanceXAuthRefreshToken, versionInstanceID, versionInstanceURL, versionInstanceCrn, versionInstanceLabel, versionInstanceCatalogID, versionInstanceOfferingID, versionInstanceKindFormat, versionInstanceVersion, versionInstanceClusterID, versionInstanceClusterRegion, versionInstanceClusterAllNamespaces)
}
