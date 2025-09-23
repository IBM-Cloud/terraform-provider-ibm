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

func TestAccIBMCrRetentionPolicyAllArgs(t *testing.T) {
	var conf containerregistryv1.RetentionPolicy
	namespace := fmt.Sprintf("tf_namespace_%d", acctest.RandIntRange(10, 100))
	imagesPerRepo := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	retainUntagged := "false"
	imagesPerRepoUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	retainUntaggedUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCrRetentionPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCrRetentionPolicyConfig(namespace, imagesPerRepo, retainUntagged),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCrRetentionPolicyExists("ibm_cr_retention_policy.cr_retention_policy", conf),
					resource.TestCheckResourceAttr("ibm_cr_retention_policy.cr_retention_policy", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_cr_retention_policy.cr_retention_policy", "images_per_repo", imagesPerRepo),
					resource.TestCheckResourceAttr("ibm_cr_retention_policy.cr_retention_policy", "retain_untagged", retainUntagged),
				),
			},
			{
				Config: testAccCheckIBMCrRetentionPolicyConfig(namespace, imagesPerRepoUpdate, retainUntaggedUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cr_retention_policy.cr_retention_policy", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_cr_retention_policy.cr_retention_policy", "images_per_repo", imagesPerRepoUpdate),
					resource.TestCheckResourceAttr("ibm_cr_retention_policy.cr_retention_policy", "retain_untagged", retainUntaggedUpdate),
				),
			},
		},
	})
}

func testAccCheckIBMCrRetentionPolicyConfig(namespace string, imagesPerRepo string, retainUntagged string) string {
	return fmt.Sprintf(`

		resource "ibm_cr_namespace" "cr_namespace" {
			name = "%s"
		}

		resource "ibm_cr_retention_policy" "cr_retention_policy" {
			namespace = ibm_cr_namespace.cr_namespace.name
			images_per_repo = %s
			retain_untagged = %s
		}
	`, namespace, imagesPerRepo, retainUntagged)
}

func testAccCheckIBMCrRetentionPolicyExists(n string, obj containerregistryv1.RetentionPolicy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		containerRegistryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContainerRegistryV1()
		if err != nil {
			return err
		}

		getRetentionPolicyOptions := &containerregistryv1.GetRetentionPolicyOptions{}

		getRetentionPolicyOptions.SetNamespace(rs.Primary.ID)

		retentionPolicy, _, err := containerRegistryClient.GetRetentionPolicy(getRetentionPolicyOptions)
		if err != nil {
			return err
		}

		obj = *retentionPolicy
		return nil
	}
}

func testAccCheckIBMCrRetentionPolicyDestroy(s *terraform.State) error {
	containerRegistryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContainerRegistryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cr_retention_policy" {
			continue
		}

		getRetentionPolicyOptions := &containerregistryv1.GetRetentionPolicyOptions{}

		getRetentionPolicyOptions.SetNamespace(rs.Primary.ID)

		// Try to find the key
		_, response, err := containerRegistryClient.GetRetentionPolicy(getRetentionPolicyOptions)

		if err == nil {
			return fmt.Errorf("cr_retention_policy still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 403 { // getRetentionPolicy returns 403 if the namespace doesn't exist
			return fmt.Errorf("[ERROR] Error checking for cr_retention_policy (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
