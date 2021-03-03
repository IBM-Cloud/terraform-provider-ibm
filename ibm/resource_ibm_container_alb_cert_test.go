// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMContainerALBCert_Basic(t *testing.T) {
	clusterName := fmt.Sprintf("tf-container-alb-%d", acctest.RandIntRange(10, 100))
	secretName := fmt.Sprintf("tf-container-alb-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerALBCertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerALBCertBasic(clusterName, secretName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_alb_cert.cert", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_alb_cert.cert", "cert_crn", certCRN),
					resource.TestCheckResourceAttr(
						"ibm_container_alb_cert.cert", "namespace", "ibm-cert-store"),
				),
			},
			{
				Config: testAccCheckIBMContainerALBCertUpdate(clusterName, secretName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_alb_cert.cert", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_alb_cert.cert", "cert_crn", updatedCertCRN),
					resource.TestCheckResourceAttr(
						"ibm_container_alb_cert.cert", "namespace", "ibm-cert-store"),
				),
			},
			{
				ResourceName:            "ibm_container_alb_cert.cert",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "issuer_name"},
			},
		},
	})
}

func testAccCheckIBMContainerALBCertDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_alb_cert" {
			continue
		}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		clusterID := parts[0]
		secretName := parts[1]
		namespace := "ibm-cert-store"
		if len(parts) > 2 && len(parts[2]) > 0 {
			namespace = parts[2]
		}
		ingressClient, err := testAccProvider.Meta().(ClientSession).VpcContainerAPI()
		if err != nil {
			return err
		}

		ingressAPI := ingressClient.Ingresses()
		resp, err := ingressAPI.GetIngressSecret(clusterID, secretName, namespace)
		if err == nil && &resp != nil && resp.Status == "deleted" {
			return nil
		} else if err == nil || !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMContainerALBCertBasic(clusterName, secretName string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name              = "%s"
  datacenter        = "%s"
  default_pool_size = 1
  machine_type      = "%s"
  hardware          = "shared"
  public_vlan_id    = "%s"
  private_vlan_id   = "%s"
  wait_till       = "MasterNodeReady"
}

resource "ibm_container_alb_cert" "cert" {
  cert_crn    = "%s"
  secret_name = "%s"
  cluster_id  = ibm_container_cluster.testacc_cluster.id
  namespace = "ibm-cert-store"
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, certCRN, secretName)
}

func testAccCheckIBMContainerALBCertUpdate(clusterName, secretName string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name              = "%s"
  datacenter        = "%s"
  default_pool_size = 1
  machine_type      = "%s"
  hardware          = "shared"
  public_vlan_id    = "%s"
  private_vlan_id   = "%s"
  wait_till       = "MasterNodeReady"
}

resource "ibm_container_alb_cert" "cert" {
  cert_crn    = "%s"
  secret_name = "%s"
  cluster_id  = ibm_container_cluster.testacc_cluster.id
  namespace = "ibm-cert-store"
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, updatedCertCRN, secretName)
}
