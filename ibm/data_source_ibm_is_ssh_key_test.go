package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMISSSHKeyDatasource_basic(t *testing.T) {
	name1 := fmt.Sprintf("terraformdskeyuat_name_%d", acctest.RandInt())
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISSSHKeyConfig(publicKey, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_ssh_key.ds_key", "name", name1),
				),
			},
		},
	})
}

func testDSCheckIBMISSSHKeyConfig(publicKey, name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ssh_key" "key" {
			name = "%s"
			public_key = "%s"
		}
		data "ibm_is_ssh_key" "ds_key" {
		    name = "${ibm_is_ssh_key.key.name}"
		}`, name, publicKey)
}
