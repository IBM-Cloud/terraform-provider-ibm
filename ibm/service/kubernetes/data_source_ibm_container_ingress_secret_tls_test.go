package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerIngressSecretTLSDatasourceBasic(t *testing.T) {
	secretName := fmt.Sprintf("tf-container-secret-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerIngressSecretTLSDataSourceConfig(secretName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_ingress_secret_tls.test_ds_secret", "cert_crn"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerIngressSecretTLSDataSourceConfig(secretName string) string {
	return fmt.Sprintf(`
	resource "ibm_container_ingress_secret_tls" "test_acc_secret" {
		secret_name    = "%s"
		secret_namespace = "ibm-cert-store"
		cert_crn = "%s"
		persistence = "%t"
		cluster  = "%s"
	}
	data "ibm_container_ingress_secret_tls" "test_ds_secret" {
	  secret_name = ibm_container_ingress_secret_tls.test_acc_secret.secret_name
	  secret_namespace = ibm_container_ingress_secret_tls.test_acc_secret.secret_namespace
	  cluster = "%s"
	}`, secretName, acc.CertCRN, true, acc.ClusterName, acc.ClusterName)
}
