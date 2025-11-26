package vpc_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test SSH Key data
const (
	rsaPublicKey     = `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR`
	ed25519PublicKey = `ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAILw6jUMh6BHClr2dUV1LHHKBQQKyTDZvY/0BSPDQmzWo`
)

func TestAccIBMIsSSHKeyNew_RSA(t *testing.T) {
	var sshKeyName = fmt.Sprintf("tf-acc-test-rsa-%d", time.Now().Unix())

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMIsSSHKeyNewConfigRSA(sshKeyName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.rsa", "name", sshKeyName),
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.rsa", "type", "rsa"),
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.rsa", "public_key", strings.TrimSpace(rsaPublicKey)),
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.rsa", "tags.#", "2"),
					resource.TestCheckResourceAttrSet("ibm_is_ssh_key_new.rsa", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_ssh_key_new.rsa", "fingerprint"),
					resource.TestCheckResourceAttrSet("ibm_is_ssh_key_new.rsa", "length"),
					resource.TestCheckResourceAttrSet("ibm_is_ssh_key_new.rsa", "crn"),
				),
			},
			{
				Config: testAccIBMIsSSHKeyNewConfigRSAUpdate(sshKeyName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.rsa", "name", sshKeyName+"-updated"),
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.rsa", "tags.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMIsSSHKeyNew_ED25519(t *testing.T) {
	var sshKeyName = fmt.Sprintf("tf-acc-test-ed25519-%d", time.Now().Unix())

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMIsSSHKeyNewConfigED25519(sshKeyName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.ed25519", "name", sshKeyName),
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.ed25519", "type", "ed25519"),
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.ed25519", "public_key", strings.TrimSpace(ed25519PublicKey)),
					resource.TestCheckResourceAttrSet("ibm_is_ssh_key_new.ed25519", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_ssh_key_new.ed25519", "fingerprint"),
				),
			},
			{
				ResourceName:      "ibm_is_ssh_key_new.ed25519",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// Configuration generation functions
func testAccIBMIsSSHKeyNewConfigRSA(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_ssh_key_new" "rsa" {
	name       = %q
	public_key = %q
	type       = "rsa"
	tags       = ["test:tags", "test:tags1"]
}
`, name, strings.TrimSpace(rsaPublicKey))
}

func testAccIBMIsSSHKeyNewConfigRSAUpdate(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_ssh_key_new" "rsa" {
	name       = "%s-updated"
	public_key = %q
	type       = "rsa"
	tags       = ["test:tags", "test:tags1", "test:tags2"]
}
`, name, strings.TrimSpace(rsaPublicKey))
}

func testAccIBMIsSSHKeyNewConfigED25519(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_ssh_key_new" "ed25519" {
	name       = %q
	public_key = %q
	type       = "ed25519"
}
`, name, strings.TrimSpace(ed25519PublicKey))
}
