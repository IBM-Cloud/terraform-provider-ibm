package ibm

import (
	"fmt"
	"log"
	"strings"
	"testing"

	bluemix "github.com/IBM-Bluemix/bluemix-go"
	"github.com/IBM-Bluemix/bluemix-go/session"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"

	"github.com/IBM-Bluemix/bluemix-go/api/account/accountv2"
	v1 "github.com/IBM-Bluemix/bluemix-go/api/container/containerv1"
	"github.com/IBM-Bluemix/bluemix-go/api/mccp/mccpv2"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMContainerCluster_basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerCluster_basic(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_num", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerClusterDestroy(s *terraform.State) error {
	csClient, err := testAccProvider.Meta().(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_cluster" {
			continue
		}

		targetEnv := getClusterTargetHeaderTestACC()
		// Try to find the key
		_, err := csClient.Clusters().Find(rs.Primary.ID, targetEnv)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for cluster (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func getClusterTargetHeaderTestACC() v1.ClusterTargetHeader {
	org := cfOrganization
	space := cfSpace
	c := new(bluemix.Config)
	sess, err := session.New(c)
	if err != nil {
		log.Fatal(err)
	}

	client, err := mccpv2.New(sess)

	if err != nil {
		log.Fatal(err)
	}

	orgAPI := client.Organizations()
	myorg, err := orgAPI.FindByName(org, BluemixRegion)

	if err != nil {
		log.Fatal(err)
	}

	spaceAPI := client.Spaces()
	myspace, err := spaceAPI.FindByNameInOrg(myorg.GUID, space, BluemixRegion)

	if err != nil {
		log.Fatal(err)
	}

	accClient, err := accountv2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	accountAPI := accClient.Accounts()
	myAccount, err := accountAPI.FindByOrg(myorg.GUID, c.Region)
	if err != nil {
		log.Fatal(err)
	}

	target := v1.ClusterTargetHeader{
		OrgID:     myorg.GUID,
		SpaceID:   myspace.GUID,
		AccountID: myAccount.GUID,
	}

	return target
}

func testAccCheckIBMContainerCluster_basic(clusterName string) string {
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
  }]

  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
}	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}
