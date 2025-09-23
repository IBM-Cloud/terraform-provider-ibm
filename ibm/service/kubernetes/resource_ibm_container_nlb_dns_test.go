// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIbmContainerNlbDnsBasic(t *testing.T) {
	var conf kubernetesserviceapiv1.NlbVPCListConfig
	clusterIps := "[ \"168.1.1.1\", \"168.1.1.2\", \"168.1.1.3\" ]"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmContainerNlbDnsConfigBasic(clusterIps),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmContainerNlbDnsExists("ibm_container_nlb_dns.container_nlb_dns", conf),
					resource.TestCheckResourceAttr("ibm_container_nlb_dns.container_nlb_dns", "cluster", acc.ClusterName),
					resource.TestCheckResourceAttr("ibm_container_nlb_dns.container_nlb_dns", "nlb_ips.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIbmContainerNlbDnsConfigBasic(clusterIps string) string {
	return fmt.Sprintf(`

	  data "ibm_container_nlb_dns" "dns" {
		cluster = "%s"
	  }

	  resource "ibm_container_nlb_dns" "container_nlb_dns" {
		cluster 		= data.ibm_container_nlb_dns.dns.cluster
		nlb_host	    = data.ibm_container_nlb_dns.dns.nlb_config.0.nlb_sub_domain
		nlb_ips  		= %s
	  }
	`, acc.ClusterName, clusterIps)
}

func testAccCheckIbmContainerNlbDnsExists(n string, obj kubernetesserviceapiv1.NlbVPCListConfig) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		satClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatelliteClientSession()
		if err != nil {
			return err
		}

		getNlbDNSListOptions := &kubernetesserviceapiv1.GetNlbDNSListOptions{}
		getNlbDNSListOptions.Cluster = flex.PtrToString(rs.Primary.ID)

		nlbConfigList, _, err := satClient.GetNlbDNSList(getNlbDNSListOptions)
		if err != nil {
			return err
		}

		obj = nlbConfigList[0]
		return nil
	}
}
