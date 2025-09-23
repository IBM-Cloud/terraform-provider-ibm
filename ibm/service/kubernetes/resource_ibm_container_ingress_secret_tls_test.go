// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMContainerIngressSecretTLS_Basic(t *testing.T) {
	secretName := fmt.Sprintf("tf-container-ingress-secret-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerIngressSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerIngressSecretTLSBasic(secretName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "cluster", acc.ClusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "secret_namespace", "ibm-cert-store"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "cert_crn", acc.CertCRN),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "persistence", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "type", "TLS"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "user_managed", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "status", "created"),
				),
			},
			{
				Config: testAccCheckIBMContainerIngressSecretTLSUpdate(secretName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "secret_namespace", "ibm-cert-store"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "cert_crn", acc.UpdatedCertCRN),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "user_managed", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "status", "created"),
				),
			},
			{
				ResourceName:            "ibm_container_ingress_secret_tls.secret",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "issuer_name", "update_secret"},
			},
		},
	})
}

func TestAccIBMContainerIngressSecretTLS_InvalidName(t *testing.T) {
	secretName := fmt.Sprintf("-)tf-container-ingress-secret-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerIngressSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMContainerIngressSecretTLSBasic(secretName),
				ExpectError: regexp.MustCompile(".*should match regexp"),
			},
		},
	})
}

// test ability to flip update_secret field to get upstream secret update from secrets manager instance via ingress API
func TestAccIBMContainerIngressSecretTLS_BasicForceUpdate(t *testing.T) {
	secretName := fmt.Sprintf("tf-container-ingress-secret-name-%d", acctest.RandIntRange(10, 100))

	var originalTS string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerIngressSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerIngressSecretTLSBasic(secretName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "cluster", acc.ClusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "secret_namespace", "ibm-cert-store"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "cert_crn", acc.CertCRN),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "persistence", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "type", "TLS"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "user_managed", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "status", "created"),

					resource.TestCheckResourceAttrWith("ibm_container_ingress_secret_tls.secret", "last_updated_timestamp", func(value string) error {
						originalTS = value
						return nil
					}),
				),
			},
			{
				Config: testAccCheckIBMContainerIngressSecretTLSForceUpdate(secretName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "secret_namespace", "ibm-cert-store"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "cert_crn", acc.CertCRN),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "user_managed", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "status", "created"),
					resource.TestCheckResourceAttrWith("ibm_container_ingress_secret_tls.secret", "last_updated_timestamp", func(value string) error {
						if originalTS == value {
							return fmt.Errorf("error timestamp not changed, indicates update didnt go through. original: %s, actual: %s", originalTS, value)
						}
						originalTS = value // check if another update will execute without a change to `update_secret`
						return nil
					}),
				),
			},
			{
				Config:             testAccCheckIBMContainerIngressSecretTLSForceUpdate(secretName),
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "secret_namespace", "ibm-cert-store"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "cert_crn", acc.CertCRN),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "user_managed", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_secret_tls.secret", "status", "created"),
					resource.TestCheckResourceAttrWith("ibm_container_ingress_secret_tls.secret", "last_updated_timestamp", func(value string) error {
						if originalTS != value {
							return fmt.Errorf("error timestamp has changed, indicates update was called again even though no modification of fields. exptected: %s, actual: %s", originalTS, value)
						}
						return nil
					}),
				),
			},
			{
				ResourceName:            "ibm_container_ingress_secret_tls.secret",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "issuer_name", "update_secret"},
			},
		},
	})
}

func testAccCheckIBMContainerIngressSecretDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_ingress_secret_tls" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		clusterID := parts[0]
		secretName := parts[1]
		namespace := parts[2]

		ingressClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
		if err != nil {
			return err
		}

		ingressAPI := ingressClient.Ingresses()
		resp, err := ingressAPI.GetIngressSecret(clusterID, secretName, namespace)
		if err == nil && &resp != nil && resp.Status == "deleted" {
			return nil
		} else if err == nil || !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error checking if secret (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMContainerIngressSecretTLSBasic(secretName string) string {
	return fmt.Sprintf(`
resource "ibm_container_ingress_secret_tls" "secret" {
  cluster  = "%s"
  secret_name = "%s"
  secret_namespace = "%s"
  cert_crn    = "%s"
  persistence = "%t"
}`, acc.ClusterName, secretName, "ibm-cert-store", acc.CertCRN, true)
}

func testAccCheckIBMContainerIngressSecretTLSUpdate(secretName string) string {
	return fmt.Sprintf(`
resource "ibm_container_ingress_secret_tls" "secret" {
  cluster  = "%s"
  secret_name = "%s"
  secret_namespace = "%s"
  cert_crn    = "%s"
  persistence = "%t"
}`, acc.ClusterName, secretName, "ibm-cert-store", acc.UpdatedCertCRN, true)
}

func testAccCheckIBMContainerIngressSecretTLSForceUpdate(secretName string) string {
	return fmt.Sprintf(`
resource "ibm_container_ingress_secret_tls" "secret" {
  cluster  = "%s"
  secret_name = "%s"
  secret_namespace = "%s"
  cert_crn    = "%s"
  persistence = "%t"
  update_secret = "%d"
}`, acc.ClusterName, secretName, "ibm-cert-store", acc.CertCRN, true, 1)
}
