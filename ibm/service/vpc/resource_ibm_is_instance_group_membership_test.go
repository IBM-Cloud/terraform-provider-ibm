// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIbmIsInstanceGroupMembershipResource_Basic(t *testing.T) {
	randInt := acctest.RandIntRange(600, 700)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDQ+WiiUR1Jg3oGSmB/2//GJ3XnotriBiGN6t3iwGces6sUsvRkza1t0Mf05DKZxC/zp0WvDTvbit2gTkF9sD37OZSn5aCJk1F5URk/JNPmz25ZogkICFL4OUfhrE3mnyKio6Bk1JIEIypR5PRtGxY9vFDUfruADDLfRi+dGwHF6U9RpvrDRo3FNtI8T0GwvWwFE7bg63vLz65CjYY5XqH9z/YWz/asH6BKumkwiphLGhuGn03+DV6DkIZqr3Oh13UDjMnTdgv1y/Kou5UM3CK1dVsmLRXPEf2KUWUq1EwRfrJXkPOrBwn8to+Yydo57FgrRM9Qw8uzvKmnVxfKW6iG3oSGA0L6ROuCq1lq0MD8ySLd56+d1ftSDaUq+0/Yt9vK3olzVP0/iZobD7chbGqTLMCzL4/CaIUR/UmX08EA0Oh0DdyAdj3UUNETAj3W8gBrV6xLR7fZAJ8roX2BKb4K8Ed3YqzgiY0zgjqvpBYl9xZl0jgVX0qMFaEa6+CeGI8= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	membershipName := fmt.Sprintf("testmembershipname%d", randInt)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsInstanceGroupMembershipResourceConfigBasic(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, membershipName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group_membership.is_instance_group_membership", "name", membershipName),
				),
			},
		},
	})
}

func testAccCheckIbmIsInstanceGroupMembershipResourceConfigBasic(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, membershipName string) string {

	return testAccCheckIbmIsInstanceGroupMembershipsDataSourceConfigBasic(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName) + fmt.Sprintf(`
	resource "ibm_is_instance_group_membership" "is_instance_group_membership" {
		instance_group = ibm_is_instance_group.instance_group.id
		instance_group_membership = data.ibm_is_instance_group_memberships.is_instance_group_memberships.memberships.0.instance_group_membership
		name = "%s"
	}
	`, membershipName)
}
