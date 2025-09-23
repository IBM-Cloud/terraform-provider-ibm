// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package registry_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/container-registry-go-sdk/containerregistryv1"
)

func TestAccIBMCrNamespaceBasic(t *testing.T) {
	var conf containerregistryv1.NamespaceDetails
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCrNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCrNamespaceConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCrNamespaceExists("ibm_cr_namespace.cr_namespace", conf),
					resource.TestCheckResourceAttr("ibm_cr_namespace.cr_namespace", "name", name),
				),
			},
			{
				Config: testAccCheckIBMCrNamespaceConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cr_namespace.cr_namespace", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIBMCrNamespaceAllArgs(t *testing.T) {
	var conf containerregistryv1.NamespaceDetails
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	tags := "[ \"tag1\", \"tag2\" ]"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	tagsUpdate := "[ \"tag3\" ]"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCrNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCrNamespaceConfig(name, tags),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCrNamespaceExists("ibm_cr_namespace.cr_namespace", conf),
					resource.TestCheckResourceAttr("ibm_cr_namespace.cr_namespace", "name", name),
					resource.TestCheckResourceAttr("ibm_cr_namespace.cr_namespace", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMCrNamespaceConfig(nameUpdate, tagsUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cr_namespace.cr_namespace", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_cr_namespace.cr_namespace", "tags.#", "1"),
				),
			},
			{
				ResourceName:            "ibm_cr_namespace.cr_namespace",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"tags"},
			},
		},
	})
}

func testAccCheckIBMCrNamespaceConfigBasic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_cr_namespace" "cr_namespace" {
			name = "%s"
		}
	`, name)
}

func testAccCheckIBMCrNamespaceConfig(name string, tags string) string {
	return fmt.Sprintf(`

		data "ibm_resource_group" "default_group" {
			is_default = "true"
		}

		resource "ibm_cr_namespace" "cr_namespace" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.default_group.id
			tags = %s
		}
	`, name, tags)
}

func testAccCheckIBMCrNamespaceExists(n string, obj containerregistryv1.NamespaceDetails) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		containerRegistryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContainerRegistryV1()
		if err != nil {
			return err
		}

		listNamespaceDetailsOptions := &containerregistryv1.ListNamespaceDetailsOptions{}

		namespaceDetailsList, _, err := containerRegistryClient.ListNamespaceDetails(listNamespaceDetailsOptions)
		if err != nil {
			return err
		}

		var namespaceDetails containerregistryv1.NamespaceDetails
		for _, namespaceDetails = range namespaceDetailsList {
			if namespaceDetails.Name != nil && *namespaceDetails.Name == rs.Primary.ID {
				obj = namespaceDetails
				return nil
			}
		}

		return fmt.Errorf("Not found")
	}
}

func testAccCheckIBMCrNamespaceDestroy(s *terraform.State) error {
	containerRegistryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContainerRegistryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cr_namespace" {
			continue
		}

		listNamespaceDetailsOptions := &containerregistryv1.ListNamespaceDetailsOptions{}

		// Try to find the key
		namespaceDetailsList, response, err := containerRegistryClient.ListNamespaceDetails(listNamespaceDetailsOptions)

		if err == nil {
			var namespaceDetails containerregistryv1.NamespaceDetails
			for _, namespaceDetails = range namespaceDetailsList {
				if *namespaceDetails.Name == rs.Primary.ID {
					break
				}
			}
			if namespaceDetails.Name != nil && *namespaceDetails.Name == rs.Primary.ID {
				return fmt.Errorf("Details of a namespace. still exists: %s", rs.Primary.ID)
			}
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for cr_namespace (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
