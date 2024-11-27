package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMContainerIngressSecretOpaqueDatasourceBasic(t *testing.T) {
	secretName := fmt.Sprintf("tf-container-secret-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerIngressSecretOpaqueDataSourceConfig(secretName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_container_ingress_secret_opaque.test_ds_secret", "fields.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerIngressSecretOpaqueDataSourceConfig(secretName string) string {
	return fmt.Sprintf(`
	resource "ibm_container_ingress_secret_opaque" "test_acc_secret" {
		secret_name    = "%s"
		secret_namespace = "ibm-cert-store"
		persistence = "%t"
		cluster  = "%s"
		fields {
			crn = "%s"
		}
	}
	data "ibm_container_ingress_secret_opaque" "test_ds_secret" {
	  secret_name = ibm_container_ingress_secret_opaque.test_acc_secret.secret_name
	  secret_namespace = ibm_container_ingress_secret_opaque.test_acc_secret.secret_namespace
	  cluster = "%s"
	}`, secretName, true, acc.ClusterName, acc.SecretCRN, acc.ClusterName)
}
