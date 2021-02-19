/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

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
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	secretName := fmt.Sprintf("terraform-secret%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerALBCertDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerALBCertBasic(clusterName, secretName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_alb_cert.cert", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_alb_cert.cert", "cert_crn", certCRN),
				),
			},
			{
				Config: testAccCheckIBMContainerALBCertUpdate(clusterName, secretName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_alb_cert.cert", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_alb_cert.cert", "cert_crn", updatedCertCRN),
				),
			},
		},
	})
}
func TestAccIBMContainerALBCert_Namespace(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	secretName := fmt.Sprintf("terraform-secret%d", acctest.RandIntRange(10, 100))
	namespaceName := "ibm-cert-store"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerALBCertDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerALBCertNameSpace(clusterName, secretName, namespaceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_alb_cert.cert", "secret_name", secretName),
					resource.TestCheckResourceAttr(
						"ibm_container_alb_cert.cert", "cert_crn", certCRN),
				),
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
}

resource "ibm_container_alb_cert" "cert" {
  cert_crn    = "%s"
  secret_name = "%s"
  cluster_id  = ibm_container_cluster.testacc_cluster.id
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, certCRN, secretName)
}

func testAccCheckIBMContainerALBCertNameSpace(clusterName, secretName, namespaceName string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name              = "%s"
  datacenter        = "%s"
  default_pool_size = 1
  machine_type      = "%s"
  hardware          = "shared"
  public_vlan_id    = "%s"
  private_vlan_id   = "%s"
}

resource "ibm_container_alb_cert" "cert" {
  cert_crn    = "%s"
  secret_name = "%s"
  cluster_id  = ibm_container_cluster.testacc_cluster.id
  namespace = "%s"
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, certCRN, secretName, namespaceName)
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
}

resource "ibm_container_alb_cert" "cert" {
  cert_crn    = "%s"
  secret_name = "%s"
  cluster_id  = ibm_container_cluster.testacc_cluster.id
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, updatedCertCRN, secretName)
}
