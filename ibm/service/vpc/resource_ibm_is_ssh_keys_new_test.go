package vpc_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test SSH Key data
const (
	rsaPublicKey     = `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR \n`
	ed25519PublicKey = `ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAILw6jUMh6BHClr2dUV1LHHKBQQKyTDZvY/0BSPDQmzWo \n`
)

func TestAccIBMIsSSHKey_RSA(t *testing.T) {
	var sshKeyName = fmt.Sprintf("tf-acc-test-rsa-%d", time.Now().Unix())

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccFrameworkPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMIsSSHKeyConfigRSA(sshKeyName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.rsa", "name", sshKeyName),
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.rsa", "type", "rsa"),
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.rsa", "public_key", strings.TrimSpace(rsaPublicKey)),
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.rsa", "tags.#", "2"),
					resource.TestCheckResourceAttrSet("ibm_is_ssh_key_new.rsa", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_ssh_key_new.rsa", "fingerprint"),
				),
			},
			{
				ResourceName:      "ibm_is_ssh_key_new.rsa",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIBMIsSSHKeyConfigRSAUpdate(sshKeyName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.rsa", "name", sshKeyName+"-updated"),
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.rsa", "tags.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMIsSSHKey_ED25519(t *testing.T) {
	var sshKeyName = fmt.Sprintf("tf-acc-test-ed25519-%d", time.Now().Unix())

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccFrameworkPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMIsSSHKeyConfigED25519(sshKeyName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.ed25519", "name", sshKeyName),
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.ed25519", "type", "ed25519"),
					resource.TestCheckResourceAttr("ibm_is_ssh_key_new.ed25519", "public_key", strings.TrimSpace(ed25519PublicKey)),
					resource.TestCheckResourceAttrSet("ibm_is_ssh_key_new.ed25519", "id"),
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
func testAccIBMIsSSHKeyConfigRSA(name string) string {
	return acctest.ConfigCompose(
		acctest.ConfigureProvider(acctest.Region()),
		fmt.Sprintf(`
			resource "ibm_is_ssh_key_new" "rsa" {
				name       = %q
				public_key = %q
				type       = "rsa"
				tags       = ["test:tags", "test:tags1"]
			}
`, name, rsaPublicKey))
}

func testAccIBMIsSSHKeyConfigRSAUpdate(name string) string {
	return acctest.ConfigCompose(
		acctest.ConfigureProvider(acctest.Region()),
		fmt.Sprintf(`
resource "ibm_is_ssh_key_new" "rsa" {
  name       = "%s-updated"
  public_key = %q
  type       = "rsa"
  tags       = ["test:tags", "test:tags1", "test:tags2"]
}
`, name, strings.TrimSpace(rsaPublicKey)))
}

func testAccIBMIsSSHKeyConfigED25519(name string) string {
	return acctest.ConfigCompose(
		acctest.ConfigureProvider(acctest.Region()),
		fmt.Sprintf(`
resource "ibm_is_ssh_key_new" "ed25519" {
  name       = %q
  public_key = %q
  type       = "ed25519"
  tags       = ["test:tags", "test:tags1"]
}
`, name, strings.TrimSpace(ed25519PublicKey)))
}
