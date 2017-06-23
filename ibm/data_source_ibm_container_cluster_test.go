package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMContainerClusterDataSource_basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	serviceKeyName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerClusterDataSource(clusterName, serviceName, serviceKeyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_cluster.testacc_ds_cluster", "worker_count", "1"),
					resource.TestCheckResourceAttr("data.ibm_container_cluster.testacc_ds_cluster", "bounded_services.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerClusterDataSource(clusterName, serviceName, serviceKeyName string) string {
	return fmt.Sprintf(`
data "ibm_org" "testacc_ds_org" {
    org = "%s"
}
data "ibm_space" "testacc_ds_space" {
    org = "%s"
    space = "%s"
}
data "ibm_account" "testacc_acc" {
    org_guid = "${data.ibm_org.testacc_ds_org.id}"
}
resource "ibm_container_cluster" "testacc_cluster" {
    name = "%s"
    datacenter = "dal10"
    org_guid = "${data.ibm_org.testacc_ds_org.id}"
    space_guid = "${data.ibm_space.testacc_ds_space.id}"
    account_guid = "${data.ibm_account.testacc_acc.id}"
   workers = [{
    name = "worker1"
    action = "add"
  }]
    machine_type = "free"
    isolation = "public"
    public_vlan_id = "vlan"
    private_vlan_id = "vlan"
}
resource "ibm_service_instance" "service" {
  name       = "%s"
  space_guid = "${data.ibm_space.testacc_ds_space.id}"
  service    = "cloudantNoSQLDB"
  plan       = "Lite"
  tags       = ["cluster-service", "cluster-bind"]
}
resource "ibm_service_key" "serviceKey" {
    name = "%s"
    service_instance_guid = "${ibm_service_instance.service.id}"
}
resource "ibm_container_bind_service" "bind_service" {
  cluster_name_id          = "${ibm_container_cluster.testacc_cluster.name}"
  service_instance_space_guid              = "${data.ibm_space.testacc_ds_space.id}"
  service_instance_name_id = "${ibm_service_instance.service.id}"
  namespace_id             = "default"
  org_guid = "${data.ibm_org.testacc_ds_org.id}"
    space_guid = "${data.ibm_space.testacc_ds_space.id}"
    account_guid = "${data.ibm_account.testacc_acc.id}"
}
data "ibm_container_cluster" "testacc_ds_cluster" {
    org_guid = "${data.ibm_org.testacc_ds_org.id}"
    space_guid = "${data.ibm_space.testacc_ds_space.id}"
    account_guid = "${data.ibm_account.testacc_acc.id}"
    cluster_name_id = "${ibm_container_cluster.testacc_cluster.id}"
}
`, cfOrganization, cfOrganization, cfSpace, clusterName, serviceName, serviceKeyName)
}
