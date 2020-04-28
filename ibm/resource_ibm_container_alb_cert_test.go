package ibm

import (
	"fmt"
	"strings"
	"testing"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
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
		targetEnv := v1.ClusterTargetHeader{
			Region: "us-south",
		}

		csClient, err := testAccProvider.Meta().(ClientSession).ContainerAPI()
		if err != nil {
			return err
		}
		albAPI := csClient.Albs()
		_, err = albAPI.GetClusterALBCertBySecretName(clusterID, secretName, targetEnv)

		if err != nil && !strings.Contains(err.Error(), "404") {
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
  region      = "%s"
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, certCRN, secretName, csRegion)
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
  region      = "%s"
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, updatedCertCRN, secretName, csRegion)
}
