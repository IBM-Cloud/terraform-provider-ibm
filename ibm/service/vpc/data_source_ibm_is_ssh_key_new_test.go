package vpc_test

import (
	"fmt"
	"log"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test SSH Key data - constants shared with resource_ibm_is_ssh_key_new_test.go

func TestAccVPCIsSshKeyNewDatasource(t *testing.T) {
	log.Printf("[INFO] UJJK Inside test file")
	dataSourceAddress := "data.ibm_is_ssh_key_new.this"
	name1 := fmt.Sprintf("tfssh-name-%d", 10)
	// 	publicKey := strings.TrimSpace(`
	// ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
	// `)
	keytype := "ed25519"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acc.ProtoV6ProviderFactories,
		// PreCheck:                 func() { acc.TestAccPreCheck(t) },
		// CheckDestroy:             testAccCheckGroupDestroy(t, groupName),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCIsSshKeyDatasourceConfig(name1, ed25519PublicKey, keytype),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceAddress, "id"),
					resource.TestCheckResourceAttrSet(dataSourceAddress, "href"),
					resource.TestCheckResourceAttrSet(dataSourceAddress, "crn"),
				),
			},
		},
	})
}

func testAccVPCIsSshKeyDatasourceConfig(name, publickey, keytype string) string {
	return fmt.Sprintf(`
	resource ibm_is_ssh_key_new ed25519 {
	  name = "%s"
	  public_key = "%s"
	  type = "%s"
	}
	data ibm_is_ssh_key_new this {
	  name = ibm_is_ssh_key_new.ed25519.name
	}
`, name, publickey, keytype)
}
