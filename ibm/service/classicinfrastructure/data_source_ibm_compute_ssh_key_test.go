// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMComputeSSHKeyDataSource_basic(t *testing.T) {
	label := fmt.Sprintf("ssh_key_test_ds_label_%d", acctest.RandIntRange(10, 100))
	notes := fmt.Sprintf("ssh_key_test_ds_notes_%d", acctest.RandIntRange(10, 100))

	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeSSHKeyDataSourceConfig(label, notes, publicKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_compute_ssh_key.testacc_ds_ssh_key", "public_key", publicKey),
					resource.TestCheckResourceAttr("data.ibm_compute_ssh_key.testacc_ds_ssh_key", "notes", notes),
					resource.TestMatchResourceAttr("data.ibm_compute_ssh_key.testacc_ds_ssh_key", "fingerprint", regexp.MustCompile("^[SHA256]")),
				),
			},
		},
	})
}

func testAccCheckIBMComputeSSHKeyDataSourceConfig(label, notes, publicKey string) string {
	return fmt.Sprintf(`
resource "ibm_compute_ssh_key" "testacc_ds_ssh_key" {
    label = "%s"
    notes = "%s"
    public_key = "%s"
}
data "ibm_compute_ssh_key" "testacc_ds_ssh_key" {
    label = "${ibm_compute_ssh_key.testacc_ds_ssh_key.label}"
}`, label, notes, publicKey)
}
