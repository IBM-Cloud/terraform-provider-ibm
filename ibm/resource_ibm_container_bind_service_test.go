package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMContainerBindService_basic(t *testing.T) {

	serviceName := fmt.Sprintf("terraform-%d", acctest.RandInt())
	serviceKey := fmt.Sprintf("terraform_%d", acctest.RandInt())
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerBindService_basic(clusterName, serviceName, serviceKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_bind_service.bind_service", "namespace_id", "default"),
				),
			},
		},
	})
}

func TestAccIBMContainerBindService_withTag(t *testing.T) {

	serviceName := fmt.Sprintf("terraform-%d", acctest.RandInt())
	serviceKey := fmt.Sprintf("terraform_%d", acctest.RandInt())
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerBindServiceWithTag(clusterName, serviceName, serviceKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_container_bind_service.bind_service", "namespace_id", "default"),
					resource.TestCheckResourceAttr("ibm_container_bind_service.bind_service", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_container_bind_service.bind_service", "cluster_name_id", clusterName),
				)},
		},
	})
}

func TestAccIBMContainerBindService_WithoutOptionalFields(t *testing.T) {

	serviceName := fmt.Sprintf("terraform-%d", acctest.RandInt())
	serviceKey := fmt.Sprintf("terraform_%d", acctest.RandInt())
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerBindService_WithoutOptionalFields(clusterName, serviceName, serviceKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_bind_service.bind_service", "namespace_id", "default"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerBindService_WithoutOptionalFields(clusterName, serviceName, serviceKey string) string {
	return fmt.Sprintf(`
data "ibm_org" "org" {
    org = "%s"
}
data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}
data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  account_guid = "${data.ibm_account.acc.id}"

  workers = [{
    name = "worker1"

    action = "add"
  }]

  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
}
resource "ibm_service_instance" "service" {
  name       = "%s"
  space_guid = "${data.ibm_space.space.id}"
  service    = "cloudantNoSQLDB"
  plan       = "Lite"
  tags       = ["cluster-service", "cluster-bind"]
}

resource "ibm_service_key" "serviceKey" {
	name = "%s"
	service_instance_guid = "${ibm_service_instance.service.id}"
}
resource "ibm_container_bind_service" "bind_service" {
  cluster_name_id          = "${ibm_container_cluster.testacc_cluster.id}"
  service_instance_space_guid              = "${data.ibm_space.space.id}"
  service_instance_name_id = "${ibm_service_instance.service.id}"
  namespace_id 			   = "default"
  account_guid = "${data.ibm_account.acc.id}"
}
	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID, serviceName, serviceKey)
}

func testAccCheckIBMContainerBindService_basic(clusterName, serviceName, serviceKey string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"

  workers = [{
    name = "worker1"

    action = "add"
  }]

  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
}


resource "ibm_service_instance" "service" {
  name       = "%s"
  space_guid = "${data.ibm_space.space.id}"
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
  service_instance_space_guid              = "${data.ibm_space.space.id}"
  service_instance_name_id = "${ibm_service_instance.service.id}"
  namespace_id 			   = "default"
  org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"
}
	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID, serviceName, serviceKey)
}

func testAccCheckIBMContainerBindServiceWithTag(clusterName, serviceName, serviceKey string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"

  workers = [{
    name = "worker1"

    action = "add"
  }]

  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
}


resource "ibm_service_instance" "service" {
  name       = "%s"
  space_guid = "${data.ibm_space.space.id}"
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
  service_instance_space_guid              = "${data.ibm_space.space.id}"
  service_instance_name_id = "${ibm_service_instance.service.id}"
  namespace_id 			   = "default"
  org_guid = "${data.ibm_org.org.id}"
  space_guid = "${data.ibm_space.space.id}"
  account_guid = "${data.ibm_account.acc.id}"
  tags = ["test"]
}
	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID, serviceName, serviceKey)
}
