// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMContainerIngressSecretOpaque_Basic(t *testing.T) {
	secretName := fmt.Sprintf("tf-container-ingress-secret-name-opaque-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerIngressSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerIngressSecretOpaqueBasic(secretName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "cluster", acc.ClusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "secret_namespace", "ibm-cert-store"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "persistence", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "type", "Opaque"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "fields.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "user_managed", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "status", "created"),
				),
			},
			{
				Config: testAccCheckIBMContainerIngressSecretOpaqueUpdate(secretName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "cluster", acc.ClusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "secret_namespace", "ibm-cert-store"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "persistence", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "type", "Opaque"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "fields.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "user_managed", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "status", "created"),
				),
			},
			{
				ResourceName:            "ibm_container_ingress_secret_opaque.secret",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "issuer_name", "update_secret"},
			},
		},
	})
}

func TestAccIBMContainerIngressSecretOpaque_InvalidName(t *testing.T) {
	secretName := fmt.Sprintf(")-tf-container-ingress-secret-name-opaque-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerIngressSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMContainerIngressSecretOpaqueBasic(secretName),
				ExpectError: regexp.MustCompile(".*should match regexp"),
			},
		},
	})
}

func TestAccIBMContainerIngressSecretOpaque_ForceUpdate(t *testing.T) {
	secretName := fmt.Sprintf("tf-container-ingress-secret-name-opaque-update-%d", acctest.RandIntRange(10, 100))

	var originalTS string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerIngressSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerIngressSecretOpaqueForceUpdateCreate(secretName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "cluster", acc.ClusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "secret_namespace", "ibm-cert-store"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "persistence", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "type", "Opaque"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "fields.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "user_managed", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "status", "created"),
					resource.TestCheckResourceAttrWith("ibm_container_ingress_secret_opaque.secret", "last_updated_timestamp", func(value string) error {
						originalTS = value
						return nil
					}),
				),
			},
			{
				Config: testAccCheckIBMContainerIngressSecretOpaqueForceUpdate(secretName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "cluster", acc.ClusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "secret_namespace", "ibm-cert-store"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "persistence", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "type", "Opaque"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "fields.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "user_managed", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "status", "created"),
					resource.TestCheckResourceAttrWith("ibm_container_ingress_secret_opaque.secret", "last_updated_timestamp", func(value string) error {
						if originalTS == value {
							return fmt.Errorf("error timestamp not changed, indicates update didnt go through. original: %s, actual: %s", originalTS, value)
						}
						originalTS = value // check if another update will execute without a change to `update_secret`
						return nil
					}),
				),
			},
			{
				Config:             testAccCheckIBMContainerIngressSecretOpaqueForceUpdate(secretName),
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "cluster", acc.ClusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "secret_namespace", "ibm-cert-store"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "persistence", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "type", "Opaque"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "fields.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "user_managed", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_opaque.secret", "status", "created"),
					resource.TestCheckResourceAttrWith("ibm_container_ingress_secret_opaque.secret", "last_updated_timestamp", func(value string) error {
						if originalTS != value {
							return fmt.Errorf("error timestamp has changed, indicates update was called again even though no modification of fields. exptected: %s, actual: %s", originalTS, value)
						}
						return nil
					}),
				),
			},
			{
				ResourceName:            "ibm_container_ingress_secret_opaque.secret",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "issuer_name", "update_secret"},
			},
		},
	})
}

func testAccCheckIBMContainerIngressSecretOpaqueDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_ingress_secret_opaque" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		clusterID := parts[0]
		secretName := parts[1]
		secretNamespace := parts[2]

		ingressClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
		if err != nil {
			return err
		}

		ingressAPI := ingressClient.Ingresses()
		resp, err := ingressAPI.GetIngressSecret(clusterID, secretName, secretNamespace)
		if err == nil && &resp != nil && resp.Status == "deleted" {
			return nil
		} else if err == nil || !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error checking if secret (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMContainerIngressSecretOpaqueBasic(secretName string) string {
	return fmt.Sprintf(`
resource "ibm_container_ingress_secret_opaque" "secret" {
  secret_name = "%s"
  secret_namespace = "%s"
  cluster  = "%s"
  persistence = "%t"
  fields {
	crn = "%s"
  }
  fields {
	crn = "%s"
  }
}`, secretName, "ibm-cert-store", acc.ClusterName, true, acc.SecretCRN, acc.SecretCRN2)
}

func testAccCheckIBMContainerIngressSecretOpaqueUpdate(secretName string) string {
	return fmt.Sprintf(`
resource "ibm_container_ingress_secret_opaque" "secret" {
  secret_name = "%s"
  secret_namespace = "%s"
  cluster  = "%s"
  persistence = "%t"
  fields {
	crn = "%s"
  }
}`, secretName, "ibm-cert-store", acc.ClusterName, true, acc.SecretCRN)
}

func testAccCheckIBMContainerIngressSecretOpaqueForceUpdateCreate(secretName string) string {
	return fmt.Sprintf(`
resource "ibm_container_ingress_secret_opaque" "secret" {
  secret_name = "%s"
  secret_namespace = "%s"
  cluster  = "%s"
  persistence = "%t"
  fields {
	crn = "%s"
  }
}`, secretName, "ibm-cert-store", acc.ClusterName, true, acc.SecretCRN)
}

func testAccCheckIBMContainerIngressSecretOpaqueForceUpdate(secretName string) string {
	return fmt.Sprintf(`
resource "ibm_container_ingress_secret_opaque" "secret" {
  secret_name = "%s"
  secret_namespace = "%s"
  cluster  = "%s"
  persistence = "%t"
  update_secret = "%d"
  fields {
	crn = "%s"
  }
}`, secretName, "ibm-cert-store", acc.ClusterName, true, 1, acc.SecretCRN)
}
