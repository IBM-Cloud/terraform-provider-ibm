// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
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
				ImportStateVerifyIgnore: []string{"region", "issuer_name"},
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
